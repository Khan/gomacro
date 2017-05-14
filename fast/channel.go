// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * channel.go
 *
 *  Created on May 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) Recv(node *ast.UnaryExpr, xe *Expr) *Expr {
	t := xe.Type
	if t.Kind() != r.Chan {
		return c.badUnaryExpr("expecting channel, found", node, xe)
	}

	if t.ChanDir()&r.RecvDir == 0 {
		return c.badUnaryExpr("cannot receive from send-only channel", node, xe)
	}

	var fun func(env *Env) (r.Value, []r.Value)
	switch x := xe.Fun.(type) {
	case func(env *Env) (r.Value, []r.Value):
		channelfun := x
		fun = func(env *Env) (r.Value, []r.Value) {
			channel, _ := channelfun(env)
			retv, ok := channel.Recv()
			var okv r.Value
			if ok {
				okv = base.True
			} else {
				okv = base.False
			}
			return retv, []r.Value{retv, okv}
		}
	default:
		channelfun := xe.AsX1()
		fun = func(env *Env) (r.Value, []r.Value) {
			retv, ok := channelfun(env).Recv()
			var okv r.Value
			if ok {
				okv = base.True
			} else {
				okv = base.False
			}
			return retv, []r.Value{retv, okv}
		}
	}
	types := []xr.Type{t.Elem(), xr.TypeOfBool}
	return exprXV(types, fun)
}
func (c *Comp) Recv1(node *ast.UnaryExpr, xe *Expr) *Expr {
	t := xe.Type
	if t.Kind() != r.Chan {
		return c.badUnaryExpr("expecting channel, found", node, xe)
	}

	if t.ChanDir()&r.RecvDir == 0 {
		return c.badUnaryExpr("cannot receive from send-only channel", node, xe)
	}

	telem := t.Elem()
	var fun I
	switch x := xe.Fun.(type) {
	case func(env *Env) (r.Value, []r.Value):
		channelfun := x
		switch telem.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return int(retv.Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return int8(retv.Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return int16(retv.Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return int32(retv.Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return uint(retv.Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return uint8(retv.Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return uint16(retv.Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return uint32(retv.Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return uintptr(retv.Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return float32(retv.Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return complex64(retv.Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.String()
			}
		default:
			fun = func(env *Env) r.Value {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv
			}

		}
	default:
		recvonly := t.ChanDir() == r.RecvDir
		channelfun := xe.AsX1()
		switch telem.Kind() {
		case r.Bool:
			if telem != BasicType(r.Bool,
			) {
				fun = func(env *Env) bool {
					retv, _ := channelfun(env).Recv()
					return retv.Bool()
				}
			} else if recvonly {
				fun = func(env *Env) bool {
					channel := channelfun(env).Interface().(<-chan bool)
					return <-channel
				}
			} else {
				fun = func(env *Env) bool {
					channel := channelfun(env).Interface().(chan bool)
					return <-channel
				}
			}
		case r.Int:
			if telem != BasicType(r.Int,
			) {
				fun = func(env *Env) int {
					retv, _ := channelfun(env).Recv()
					return int(retv.Int())
				}
			} else if recvonly {
				fun = func(env *Env) int {
					channel := channelfun(env).Interface().(<-chan int)
					return <-channel
				}
			} else {
				fun = func(env *Env) int {
					channel := channelfun(env).Interface().(chan int)
					return <-channel
				}
			}
		case r.Int8:
			if telem != BasicType(r.Int8,
			) {
				fun = func(env *Env) int8 {
					retv, _ := channelfun(env).Recv()
					return int8(retv.Int())
				}
			} else if recvonly {
				fun = func(env *Env) int8 {
					channel := channelfun(env).Interface().(<-chan int8)
					return <-channel
				}
			} else {
				fun = func(env *Env) int8 {
					channel := channelfun(env).Interface().(chan int8)
					return <-channel
				}
			}
		case r.Int16:
			if telem != BasicType(r.Int16,
			) {
				fun = func(env *Env) int16 {
					retv, _ := channelfun(env).Recv()
					return int16(retv.Int())
				}
			} else if recvonly {
				fun = func(env *Env) int16 {
					channel := channelfun(env).Interface().(<-chan int16)
					return <-channel
				}
			} else {
				fun = func(env *Env) int16 {
					channel := channelfun(env).Interface().(chan int16)
					return <-channel
				}
			}
		case r.Int32:
			if telem != BasicType(r.Int32,
			) {
				fun = func(env *Env) int32 {
					retv, _ := channelfun(env).Recv()
					return int32(retv.Int())
				}
			} else if recvonly {
				fun = func(env *Env) int32 {
					channel := channelfun(env).Interface().(<-chan int32)
					return <-channel
				}
			} else {
				fun = func(env *Env) int32 {
					channel := channelfun(env).Interface().(chan int32)
					return <-channel
				}
			}
		case r.Int64:
			if telem != BasicType(r.Int64,
			) {
				fun = func(env *Env) int64 {
					retv, _ := channelfun(env).Recv()
					return retv.Int()
				}
			} else if recvonly {
				fun = func(env *Env) int64 {
					channel := channelfun(env).Interface().(<-chan int64)
					return <-channel
				}
			} else {
				fun = func(env *Env) int64 {
					channel := channelfun(env).Interface().(chan int64)
					return <-channel
				}
			}
		case r.Uint:
			if telem != BasicType(r.Uint,
			) {
				fun = func(env *Env) uint {
					retv, _ := channelfun(env).Recv()
					return uint(retv.Uint())
				}
			} else if recvonly {
				fun = func(env *Env) uint {
					channel := channelfun(env).Interface().(<-chan uint)
					return <-channel
				}
			} else {
				fun = func(env *Env) uint {
					channel := channelfun(env).Interface().(chan uint)
					return <-channel
				}
			}
		case r.Uint8:
			if telem != BasicType(r.Uint8,
			) {
				fun = func(env *Env) uint8 {
					retv, _ := channelfun(env).Recv()
					return uint8(retv.Uint())
				}
			} else if recvonly {
				fun = func(env *Env) uint8 {
					channel := channelfun(env).Interface().(<-chan uint8)
					return <-channel
				}
			} else {
				fun = func(env *Env) uint8 {
					channel := channelfun(env).Interface().(chan uint8)
					return <-channel
				}
			}
		case r.Uint16:
			if telem != BasicType(r.Uint16,
			) {
				fun = func(env *Env) uint16 {
					retv, _ := channelfun(env).Recv()
					return uint16(retv.Uint())
				}
			} else if recvonly {
				fun = func(env *Env) uint16 {
					channel := channelfun(env).Interface().(<-chan uint16)
					return <-channel
				}
			} else {
				fun = func(env *Env) uint16 {
					channel := channelfun(env).Interface().(chan uint16)
					return <-channel
				}
			}
		case r.Uint32:
			if telem != BasicType(r.Uint32,
			) {
				fun = func(env *Env) uint32 {
					retv, _ := channelfun(env).Recv()
					return uint32(retv.Uint())
				}
			} else if recvonly {
				fun = func(env *Env) uint32 {
					channel := channelfun(env).Interface().(<-chan uint32)
					return <-channel
				}
			} else {
				fun = func(env *Env) uint32 {
					channel := channelfun(env).Interface().(chan uint32)
					return <-channel
				}
			}
		case r.Uint64:
			if telem != BasicType(r.Uint64,
			) {
				fun = func(env *Env) uint64 {
					retv, _ := channelfun(env).Recv()
					return retv.Uint()
				}
			} else if recvonly {
				fun = func(env *Env) uint64 {
					channel := channelfun(env).Interface().(<-chan uint64)
					return <-channel
				}
			} else {
				fun = func(env *Env) uint64 {
					channel := channelfun(env).Interface().(chan uint64)
					return <-channel
				}
			}
		case r.Uintptr:
			if telem != BasicType(r.Uintptr,
			) {
				fun = func(env *Env) uintptr {
					retv, _ := channelfun(env).Recv()
					return uintptr(retv.Uint())
				}
			} else if recvonly {
				fun = func(env *Env) uintptr {
					channel := channelfun(env).Interface().(<-chan uintptr)
					return <-channel
				}
			} else {
				fun = func(env *Env) uintptr {
					channel := channelfun(env).Interface().(chan uintptr)
					return <-channel
				}
			}
		case r.Float32:
			if telem != BasicType(r.Float32,
			) {
				fun = func(env *Env) float32 {
					retv, _ := channelfun(env).Recv()
					return float32(retv.Float())
				}
			} else if recvonly {
				fun = func(env *Env) float32 {
					channel := channelfun(env).Interface().(<-chan float32)
					return <-channel
				}
			} else {
				fun = func(env *Env) float32 {
					channel := channelfun(env).Interface().(chan float32)
					return <-channel
				}
			}
		case r.Float64:
			if telem != BasicType(r.Float64,
			) {
				fun = func(env *Env) float64 {
					retv, _ := channelfun(env).Recv()
					return retv.Float()
				}
			} else if recvonly {
				fun = func(env *Env) float64 {
					channel := channelfun(env).Interface().(<-chan float64)
					return <-channel
				}
			} else {
				fun = func(env *Env) float64 {
					channel := channelfun(env).Interface().(chan float64)
					return <-channel
				}
			}
		case r.Complex64:
			if telem != BasicType(r.Complex64,
			) {
				fun = func(env *Env) complex64 {
					retv, _ := channelfun(env).Recv()
					return complex64(retv.Complex())
				}
			} else if recvonly {
				fun = func(env *Env) complex64 {
					channel := channelfun(env).Interface().(<-chan complex64)
					return <-channel
				}
			} else {
				fun = func(env *Env) complex64 {
					channel := channelfun(env).Interface().(chan complex64)
					return <-channel
				}
			}
		case r.Complex128:
			if telem != BasicType(r.Complex128,
			) {
				fun = func(env *Env) complex128 {
					retv, _ := channelfun(env).Recv()
					return retv.Complex()
				}
			} else if recvonly {
				fun = func(env *Env) complex128 {
					channel := channelfun(env).Interface().(<-chan complex128)
					return <-channel
				}
			} else {
				fun = func(env *Env) complex128 {
					channel := channelfun(env).Interface().(chan complex128)
					return <-channel
				}
			}
		case r.String:
			if telem != BasicType(r.String,
			) {
				fun = func(env *Env) string {
					retv, _ := channelfun(env).Recv()
					return retv.String()
				}
			} else if recvonly {
				fun = func(env *Env) string {
					channel := channelfun(env).Interface().(<-chan string)
					return <-channel
				}
			} else {
				fun = func(env *Env) string {
					channel := channelfun(env).Interface().(chan string)
					return <-channel
				}
			}
		default:
			fun = func(env *Env) r.Value {
				retv, _ := channelfun(env).Recv()
				return retv
			}

		}
	}
	return exprFun(telem, fun)
}
func (c *Comp) Send(node *ast.SendStmt) {
	channel := c.Expr1(node.Chan)
	t := channel.Type
	if t.Kind() != r.Chan {
		c.Errorf("cannot send to non-channel type %v: %v", t, node)
		return
	}
	if t.ChanDir()&r.SendDir == 0 {
		c.Errorf("cannot send to receive-only channel type %v: %v", t, node)
		return
	}
	telem := t.Elem()
	expr := c.Expr1(node.Value)
	if expr.Const() {
		expr.ConstTo(telem)
	} else if expr.Type == nil || !expr.Type.AssignableTo(telem) {
		c.Errorf("cannot use %v <%v> as type %v in send", node.Value, expr.Type, telem)
		return
	}
	channelfun := channel.AsX1()
	sendonly := t.ChanDir() == r.SendDir
	var stmt Stmt
	if expr.Const() {
		v := r.ValueOf(expr.Value)
		switch telem {
		case xr.TypeOfBool:
			value := v.Bool()
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- bool)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan bool)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfInt:
			value :=

				int(v.Int())
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- int)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan int)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfInt8:
			value :=

				int8(v.Int())
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- int8)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan int8)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfInt16:
			value :=

				int16(v.Int())
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- int16)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan int16)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfInt32:
			value :=

				int32(v.Int())
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- int32)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan int32)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfInt64:
			value := v.Int()
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- int64)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan int64)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUint:
			value :=

				uint(v.Uint())
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uint)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uint)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUint8:
			value :=

				uint8(v.Uint())
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uint8)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uint8)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUint16:
			value :=

				uint16(v.Uint())
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uint16)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uint16)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUint32:
			value :=

				uint32(v.Uint())
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uint32)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uint32)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUint64:
			value := v.Uint()
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uint64)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uint64)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUintptr:
			value :=

				uintptr(v.Uint())
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uintptr)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uintptr)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfFloat32:
			value :=

				float32(v.Float())
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- float32)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan float32)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfFloat64:
			value := v.Float()
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- float64)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan float64)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfComplex64:
			value :=

				complex64(v.Complex())
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- complex64)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan complex64)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfComplex128:
			value := v.Complex()
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- complex128)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan complex128)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfString:
			value := v.String()
			if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- string)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan string)
					channel <- value
					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
		if stmt == nil {
			stmt = func(env *Env) (Stmt, *Env) {
				channel := channelfun(env)
				channel.Send(v)
				env.IP++
				return env.Code[env.IP], env
			}
		}

	} else {
		switch telem {
		case xr.TypeOfBool:
			if exprfun, ok := expr.Fun.(func(*Env) bool); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- bool)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan bool)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfInt:
			if exprfun, ok := expr.Fun.(func(*Env) int); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- int)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan int)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfInt8:
			if exprfun, ok := expr.Fun.(func(*Env) int8); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- int8)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan int8)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfInt16:
			if exprfun, ok := expr.Fun.(func(*Env) int16); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- int16)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan int16)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfInt32:
			if exprfun, ok := expr.Fun.(func(*Env) int32); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- int32)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan int32)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfInt64:
			if exprfun, ok := expr.Fun.(func(*Env) int64); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- int64)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan int64)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUint:
			if exprfun, ok := expr.Fun.(func(*Env) uint); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uint)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uint)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUint8:
			if exprfun, ok := expr.Fun.(func(*Env) uint8); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uint8)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uint8)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUint16:
			if exprfun, ok := expr.Fun.(func(*Env) uint16); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uint16)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uint16)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUint32:
			if exprfun, ok := expr.Fun.(func(*Env) uint32); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uint32)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uint32)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUint64:
			if exprfun, ok := expr.Fun.(func(*Env) uint64); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uint64)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uint64)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfUintptr:
			if exprfun, ok := expr.Fun.(func(*Env) uintptr); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- uintptr)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan uintptr)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfFloat32:
			if exprfun, ok := expr.Fun.(func(*Env) float32); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- float32)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan float32)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfFloat64:
			if exprfun, ok := expr.Fun.(func(*Env) float64); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- float64)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan float64)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfComplex64:
			if exprfun, ok := expr.Fun.(func(*Env) complex64); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- complex64)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan complex64)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfComplex128:
			if exprfun, ok := expr.Fun.(func(*Env) complex128); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- complex128)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan complex128)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case xr.TypeOfString:
			if exprfun, ok := expr.Fun.(func(*Env) string); !ok {
				break
			} else if sendonly {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan<- string)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				stmt = func(env *Env) (Stmt, *Env) {
					channel := channelfun(env).Interface().(chan string)
					channel <- exprfun(env)
					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
		if stmt == nil {
			exprfun := expr.AsX1()
			rtelem := telem.ReflectType()
			stmt = func(env *Env) (Stmt, *Env) {
				channel := channelfun(env)
				value := exprfun(env)
				if value.Type() != rtelem {
					value = value.Convert(rtelem)
				}

				channel.Send(value)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	}
	c.Code.Append(stmt)
}