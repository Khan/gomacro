// this file was generated by gomacro command: import _i "github.com/cosmos72/gomacro/xreflect"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package xreflect

import (
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/xreflect"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/xreflect"] = imports.Package{
		Binds: map[string]r.Value{
			"DefaultImporter":    r.ValueOf(DefaultImporter),
			"GensymEmbedded":     r.ValueOf(GensymEmbedded),
			"GensymPrivate":      r.ValueOf(GensymPrivate),
			"MaxDepth":           r.ValueOf(int64(MaxDepth)),
			"MissingMethod":      r.ValueOf(MissingMethod),
			"NewUniverse":        r.ValueOf(NewUniverse),
			"QName1":             r.ValueOf(QName1),
			"QName2":             r.ValueOf(QName2),
			"QNameGo":            r.ValueOf(QNameGo),
			"QNameGo2":           r.ValueOf(QNameGo2),
			"StrGensymEmbedded":  r.ValueOf(StrGensymEmbedded),
			"StrGensymInterface": r.ValueOf(StrGensymInterface),
			"StrGensymPrivate":   r.ValueOf(StrGensymPrivate),
			"Zero":               r.ValueOf(Zero),
		},
		Types: map[string]r.Type{
			"Error":           r.TypeOf((*Error)(nil)).Elem(),
			"Importer":        r.TypeOf((*Importer)(nil)).Elem(),
			"InterfaceHeader": r.TypeOf((*InterfaceHeader)(nil)).Elem(),
			"Method":          r.TypeOf((*Method)(nil)).Elem(),
			"Package":         r.TypeOf((*Package)(nil)).Elem(),
			"QName":           r.TypeOf((*QName)(nil)).Elem(),
			"QNameI":          r.TypeOf((*QNameI)(nil)).Elem(),
			"StructField":     r.TypeOf((*StructField)(nil)).Elem(),
			"Type":            r.TypeOf((*Type)(nil)).Elem(),
			"Types":           r.TypeOf((*Types)(nil)).Elem(),
			"Universe":        r.TypeOf((*Universe)(nil)).Elem(),
			"Value":           r.TypeOf((*Value)(nil)).Elem(),
		},
		Proxies: map[string]r.Type{
			"QNameI": r.TypeOf((*QNameI_github_com_cosmos72_gomacro_xreflect)(nil)).Elem(),
		}}
}

// --------------- proxy for github.com/cosmos72/gomacro/xreflect.QNameI ---------------
type QNameI_github_com_cosmos72_gomacro_xreflect struct {
	Object   interface{}
	Name_    func() string
	PkgPath_ func() string
}

func (Proxy *QNameI_github_com_cosmos72_gomacro_xreflect) Name() string {
	return Proxy.Name_()
}
func (Proxy *QNameI_github_com_cosmos72_gomacro_xreflect) PkgPath() string {
	return Proxy.PkgPath_()
}
