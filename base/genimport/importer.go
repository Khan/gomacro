/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * importer.go
 *
 *  Created on Feb 27, 2017
 *      Author Massimiliano Ghilardi
 */

package genimport

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/build"
	"go/types"
	"golang.org/x/mod/modfile"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	r "reflect"
	"strings"

	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/base/paths"
	"github.com/cosmos72/gomacro/base/reflect"
	"github.com/cosmos72/gomacro/imports"
)

type ImportMode int

const (
	// ImBuiltin import mechanism is:
	// 1. write a file $GOPATH/src/github.com/cosmos72/gomacro/imports/$PKGPATH.go containing a single func init()
	//    i.e. *inside* gomacro sources
	// 2. tell the user to recompile gomacro
	ImBuiltin ImportMode = iota

	// ImThirdParty import mechanism is the same as ImBuiltin, except that files are created in a thirdparty/ subdirectory:
	// 1. write a file $GOPATH/src/github.com/cosmos72/gomacro/imports/thirdparty/$PKGPATH.go containing a single func init()
	//    i.e. *inside* gomacro sources
	// 2. tell the user to recompile gomacro
	ImThirdParty

	// ImInception import mechanism is:
	// 1. write a file $GOPATH/src/$PKGPATH/x_package.go containing a single func init()
	//    i.e. *inside* the package to be imported
	// 2. tell the user to recompile $PKGPATH
	ImInception

	// ImPlugin import mechanism is:
	// 1. write a file $GOPATH/src/gomacro.imports/$PKGPATH/$PKGNAME.go containing a var Packages map[string]Package
	//    and a single func init() to populate it
	// 2. invoke "go build -buildmode=plugin" on the file to create a shared library
	// 3. load such shared library with plugin.Open().Lookup("Packages")
	ImPlugin
)

type PackageRef struct {
	imports.Package
	Path string
}

func (ref *PackageRef) DefaultName() string {
	return ref.Package.DefaultName(ref.Path)
}

func (ref *PackageRef) String() string {
	return fmt.Sprintf("{%s %q, %d binds, %d types}", ref.DefaultName(), ref.Path, len(ref.Binds), len(ref.Types))
}

type Importer struct {
	srcDir     string
	mode       types.ImportMode
	PluginOpen r.Value // = reflect.ValueOf(plugin.Open)
	output     *Output
}

func DefaultImporter(o *Output) *Importer {
	return &Importer{output: o}
}

func (imp *Importer) havePluginOpen() bool {
	if !imp.PluginOpen.IsValid() {
		imp.PluginOpen = imports.Packages["plugin"].Binds["Open"]
		if !imp.PluginOpen.IsValid() {
			imp.PluginOpen = reflect.NoneR // cache the failure
		}
	}
	return imp.PluginOpen != reflect.NoneR
}

// LookupPackage returns a package if already present in cache
func LookupPackage(alias, path string) *PackageRef {
	pkg, found := imports.Packages[path]
	if !found {
		return nil
	}
	if len(pkg.Name) == 0 {
		// missing pkg.Name, initialize it
		pkg.DefaultName(path)
		imports.Packages[path] = pkg
	}
	if len(alias) == 0 {
		// import "foo" => get alias from package name
		alias = pkg.DefaultName(path)
	}
	return &PackageRef{Package: pkg, Path: path}
}

func (imp *Importer) wrapImportError(path string, enableModule bool, err error) output.RuntimeError {
	if rerr, ok := err.(output.RuntimeError); ok {
		return rerr
	}
	if enableModule {
		return imp.output.MakeRuntimeError("error loading package %q metadata: %v", path, err)
	}
	return imp.output.MakeRuntimeError(
		"error loading package %q metadata, maybe you need to download (go get), compile (go build) and install (go install) it? %v",
		path, err)
}

func (imp *Importer) ImportPackage(alias, path string, enableModule bool) *PackageRef {
	ref, err := imp.ImportPackageOrError(alias, path, enableModule)
	if err != nil {
		panic(err)
	}
	return ref
}

func (imp *Importer) ImportPackageOrError(alias, pkgpath string, enableModule bool) (*PackageRef, error) {

	ref := LookupPackage(alias, pkgpath)
	if ref != nil {
		return ref, nil
	}
	paths.GetImportsSrcDir() // warns if GOPATH or paths.ImportsDir may be wrong

	o := imp.output
	gpkg, err := imp.Load(pkgpath, enableModule) // loads names and types, not the values!
	if err != nil {
		return nil, imp.wrapImportError(pkgpath, enableModule, err)
	}
	var mode ImportMode
	switch alias {
	case "_b":
		mode = ImBuiltin
	case "_i":
		mode = ImInception
	case "_3":
		mode = ImThirdParty
	default:
		if len(alias) == 0 {
			alias = gpkg.Name()
		}
		if imp.havePluginOpen() {
			mode = ImPlugin
		} else {
			mode = ImThirdParty
		}
	}
	file := createImportFile(imp.output, pkgpath, gpkg, mode, enableModule)
	ref = &PackageRef{Path: pkgpath}
	if len(file) == 0 || mode != ImPlugin {
		// either the package exports nothing, or user must rebuild gomacro.
		// in both cases, still cache it to avoid recreating the file.
		imports.Packages[pkgpath] = ref.Package
		return ref, nil
	}
	soname := compilePlugin(o, file, enableModule, o.Stdout, o.Stderr)
	ipkgs := imp.loadPluginSymbol(soname, "Packages")
	pkgs := *ipkgs.(*map[string]imports.PackageUnderlying)

	// cache *all* packages found for future use
	imports.Packages.Merge(pkgs)

	// but return only requested one
	pkg, found := imports.Packages[pkgpath]
	if !found {
		return nil, imp.output.MakeRuntimeError(
			"error loading package %q: the compiled plugin %q does not contain it! internal error? %v",
			pkgpath, soname)
	}
	ref.Package = pkg
	return ref, nil
}

func createImportFile(o *Output, pkgpath string, pkg *types.Package, mode ImportMode, enableModule bool) string {
	dir := computeImportDir(o, pkgpath, mode)
	if mode == ImPlugin {
		createDir(o, dir)
		removeAllFilesInDirExcept(o, dir, []string{"go.mod", "go.sum"})
	}
	f := computeImportFilename(o, pkgpath, mode)
	f = paths.Subdir(dir, f)

	buf := bytes.Buffer{}
	isEmpty := writeImportFile(o, &buf, pkgpath, pkg, mode)
	if isEmpty {
		o.Warnf("package %q exports zero constants, functions, types and variables", pkgpath)
		return ""
	}

	err := ioutil.WriteFile(f, buf.Bytes(), os.FileMode(0o644))
	if err != nil {
		o.Errorf("error writing file %q: %v", f, err)
	}
	switch mode {
	case ImBuiltin, ImThirdParty:
		o.Warnf("created file %q, recompile gomacro to use it", f)
	case ImInception:
		o.Warnf("created file %q, recompile %s to use it", f, pkgpath)
	case ImPlugin:
		// if needed, go.mod file was created already by Importer.Load()
		env := environForCompiler(enableModule)
		runGoModTidyIfNeeded(o, pkgpath, dir, env)
	}
	return f
}

func createDir(o *Output, dir string) {
	err := os.MkdirAll(dir, 0o700)
	if err != nil {
		o.Errorf("error creating directory %q: %v", dir, err)
	}
}

func listDir(o *Output, dir string) []os.FileInfo {
	d, err := os.Open(dir)
	if err != nil {
		o.Errorf("error opening directory %q: %v", dir, err)
	}
	defer d.Close()
	infos, err := d.Readdir(0)
	if err != nil {
		o.Errorf("error listing directory %q: %v", dir, err)
	}
	return infos
}

func removeAllFilesInDir(o *Output, dir string) {
	removeAllFilesInDirExcept(o, dir, nil)
}

func removeAllFilesInDirExcept(o *Output, dir string, except_list []string) {
	for _, info := range listDir(o, dir) {
		if info.IsDir() {
			continue
		}
		name := info.Name()
		for _, except_name := range except_list {
			if name == except_name {
				name = ""
				break
			}
		}
		if name == "" {
			continue
		}
		f := paths.Subdir(dir, name)
		if err := os.Remove(f); err != nil {
			o.Errorf("error removing file %q: %v", f, err)
		}
	}
}

func createPluginGoModFile(o *Output, pkgpath string, dir string) string {
	file := modfile.File{}
	err := file.AddModuleStmt("gomacro.imports/" + pkgpath)
	if err != nil {
		o.Errorf("error setting module in go.mod", err)
	}

	// Attempt to use the local module if present.
	// This only works if the import shares a mod file with current working
	// directory, because we only know to guess "." for the local location.
	// TODO: Find a way to support imports from local disk that aren't in
	//  the cwd project.
	if pkgModFileInfo, err := getModuleFileInfo("."); err == nil &&
		(pkgpath == pkgModFileInfo.Path || strings.HasPrefix(pkgpath, pkgModFileInfo.Path+"/")) {

		o.Debugf("importing %s from local %s", pkgpath, pkgModFileInfo.GoMod)
		goModReplaceDirectives(o, pkgModFileInfo, file)
	}

	gomod := paths.Subdir(dir, "go.mod")

	format, err := file.Format()
	if err != nil {
		o.Debugf("error producing go.mod %v", err)
		return ""
	}

	err = ioutil.WriteFile(gomod, format, os.FileMode(0o644))
	if err != nil {
		o.Errorf("error writing file %q: %v", gomod, err)
		return ""
	}

	return gomod
}

// goModReplaceDirectives will create the replacement directives associated
// with a module that can be found locally, for the purpose of use in a
// gomacro.imports mod file.
func goModReplaceDirectives(o *Output, pkgModFileInfo modInfo, dest modfile.File) {
	m, err := getModuleFile(pkgModFileInfo)
	if err != nil {
		o.Errorf("error getting go.mod", err)
		return
	}

	err = dest.AddReplace(pkgModFileInfo.Path, "", pkgModFileInfo.Dir, "")
	if err != nil {
		o.Debugf("error adding initial replace directive %v", err)
		return
	}

	// Copy the replace directives from the imported mod, so the functionality
	// remains similar.
	for _, replaceDirective := range m.Replace {
		newPath := replaceDirective.New.Path

		// If the replace directive is to the local disk but not absolute,
		// point to the correctly location.
		if modfile.IsDirectoryPath(replaceDirective.New.Path) &&
			!filepath.IsAbs(replaceDirective.New.Path) {

			newPath = filepath.Join(pkgModFileInfo.Dir, replaceDirective.New.Path)
		}

		err := dest.AddReplace(replaceDirective.Old.Path, replaceDirective.Old.Version,
			newPath, replaceDirective.New.Version)
		if err != nil {
			o.Debugf("error adding replace directive for %s, %v", replaceDirective.Old.String(), err)
		}
	}
}

type modInfo struct {
	Path      string `json:"Path"`
	Dir       string `json:"Dir"`
	GoMod     string `json:"GoMod"`
	GoVersion string `json:"GoVersion"`
	Main      bool   `json:"Main"`
}

func getModuleFile(i modInfo) (*modfile.File, error) {
	raw, err := ioutil.ReadFile(i.GoMod)
	if err != nil {
		return nil, fmt.Errorf("reading go.mod file: %w", err)
	}

	return modfile.Parse("go.mod", raw, nil)
}

func getModuleFileInfo(dir string) (modInfo, error) {
	// https://github.com/golang/go/issues/44753#issuecomment-790089020
	cmd := exec.Command("go", "list", "-m", "-json", "-f", "{{.GoMod}}")
	cmd.Dir = dir

	raw, err := cmd.CombinedOutput()
	if err != nil {
		return modInfo{}, fmt.Errorf("command go list: %w: %s", err, string(raw))
	}

	var v modInfo
	err = json.Unmarshal(raw, &v)
	if err != nil {
		return modInfo{}, fmt.Errorf("unmarshaling error: %w: %s", err, string(raw))
	}

	if v.GoMod == "" {
		return modInfo{}, errors.New("working directory is not part of a module")
	}
	return v, nil
}

func packageSanitizedName(path string) string {
	return sanitizeIdent(paths.FileName(path))
}

func sanitizeIdent(str string) string {
	return sanitizeIdent2(str, '_')
}

func sanitizeIdent2(str string, replacement rune) string {
	runes := []rune(str)
	for i, ch := range runes {
		if ch >= 'a' && ch <= 'z' {
			continue
		} else if ch >= 'A' && ch <= 'Z' {
			if i == 0 {
				// first rune must be lowercase to avoid conflict
				// with Packages, ValueOf, TypeOf
				runes[i] = ch - 'A' + 'a'
			}
			continue
		} else if i > 0 && (ch == '_' || (ch >= '0' && ch <= '9')) {
			continue
		}
		runes[i] = replacement
	}
	str = string(runes)
	if isReservedKeyword(str) {
		runes = append(runes, '_')
		str = string(runes)
	}
	return str
}

func computeImportDir(o *Output, pkgpath string, mode ImportMode) string {
	switch mode {
	case ImBuiltin:
		// user will need to recompile gomacro
		return paths.GetImportsSrcDir()
	case ImThirdParty:
		// either plugin.Open is not available, or user explicitly requested import _3 "package".
		// In both cases, user will need to recompile gomacro
		return paths.Subdir(paths.GetImportsSrcDir(), "thirdparty")
	case ImInception:
		// user will need to recompile the package being imported
		for _, srcdir := range paths.GoSrcDirs {
			dir := paths.Subdir(srcdir, pkgpath)
			if _, err := os.Stat(dir); err == nil {
				return paths.Subdir(srcdir, pkgpath)
			}
		}
		o.Errorf("unable to locate package %q in $GOPATH/src ($GOPATH=%s)",
			pkgpath, build.Default.GOPATH)
	case ImPlugin:
		return paths.Subdir(paths.GoSrcDir, "gomacro.imports", pkgpath)
	default:
		o.Errorf("unknown import mode: %v", mode)
	}
	return ""
}

func computeImportFilename(o *Output, pkgpath string, mode ImportMode) string {
	switch mode {
	case ImBuiltin:
		// user will need to recompile gomacro
		return sanitizeIdent(pkgpath) + ".go"
	case ImThirdParty:
		// either plugin.Open is not available, or user explicitly requested import _3 "package".
		// In both cases, user will need to recompile gomacro
		return sanitizeIdent(pkgpath) + ".go"
	case ImInception:
		// user will need to recompile package being imported
		return "x_package.go"
	case ImPlugin:
		return sanitizeIdent(paths.FileName(pkgpath)) + ".go"
	default:
		o.Errorf("unknown import mode: %v", mode)
		return ""
	}
}
