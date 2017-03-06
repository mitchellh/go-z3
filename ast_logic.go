package z3

import (
	"unsafe"
)

// #include "go-z3.h"
import "C"

// Distinct creates an AST node representing adding.
//
// All AST values must be part of the same context.
func (a *AST) Distinct(args ...*AST) *AST {
	raws := make([]C.Z3_ast, len(args)+1)
	raws[0] = a.rawAST
	for i, arg := range args {
		raws[i+1] = arg.rawAST
	}

	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_distinct(
			a.rawCtx,
			C.uint(len(raws)),
			(*C.Z3_ast)(unsafe.Pointer(&raws[0]))),
	}
}

// Not creates an AST node representing not(a)
//
// Maps to: Z3_mk_not
func (a *AST) Not() *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_not(a.rawCtx, a.rawAST),
	}
}

// Eq creates a "equal" comparison.
//
// Maps to: Z3_mk_eq
func (a *AST) Eq(a2 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_eq(a.rawCtx, a.rawAST, a2.rawAST),
	}
}

// Ite creates an AST node representing if a then a2 else a3.
//
// a and a2 must be part of the same Context and be boolean types.
func (a *AST) Ite(a2, a3 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_ite(a.rawCtx, a.rawAST, a2.rawAST, a3.rawAST),
	}
}

// Iff creates an AST node representing a iff a2.
//
// a and a2 must be part of the same Context and be boolean types.
func (a *AST) Iff(a2 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_iff(a.rawCtx, a.rawAST, a2.rawAST),
	}
}

// Implies creates an AST node representing a implies a2.
//
// a and a2 must be part of the same Context and be boolean types.
func (a *AST) Implies(a2 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_implies(a.rawCtx, a.rawAST, a2.rawAST),
	}
}

// Xor creates an AST node representing a xor a2.
//
// a and a2 must be part of the same Context and be boolean types.
func (a *AST) Xor(a2 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_xor(a.rawCtx, a.rawAST, a2.rawAST),
	}
}

// And creates an AST node representing a and a2 and ... aN.
//
// a and a2 must be part of the same Context and be boolean types.
func (a *AST) And(args ...*AST) *AST {
	raws := make([]C.Z3_ast, len(args)+1)
	raws[0] = a.rawAST
	for i, arg := range args {
		raws[i+1] = arg.rawAST
	}

	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_and(
			a.rawCtx,
			C.uint(len(raws)),
			(*C.Z3_ast)(unsafe.Pointer(&raws[0]))),
	}
}

// Or creates an AST node representing a or a2 or ... aN.
//
// a and a2 must be part of the same Context and be boolean types.
func (a *AST) Or(args ...*AST) *AST {
	raws := make([]C.Z3_ast, len(args)+1)
	raws[0] = a.rawAST
	for i, arg := range args {
		raws[i+1] = arg.rawAST
	}

	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_or(
			a.rawCtx,
			C.uint(len(raws)),
			(*C.Z3_ast)(unsafe.Pointer(&raws[0]))),
	}
}
