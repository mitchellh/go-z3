package z3

// #include "go-z3.h"
import "C"

// Model represents a model from a solver.
//
// Memory management for this is manual and based on reference counting.
// When a model is initialized (via Solver.Model for example), it always has
// a reference count of 1. You must call Close when you're done.
type Model struct {
	rawCtx   C.Z3_context
	rawModel C.Z3_model
}

// String returns a human-friendly string version of the model.
func (m *Model) String() string {
	return C.GoString(C.Z3_model_to_string(m.rawCtx, m.rawModel))
}

//-------------------------------------------------------------------
// Assignments
//-------------------------------------------------------------------

// Eval evaluates the given AST within the model. This can be used to get
// the assignment of an AST. This will return nil if evaluation failed.
//
// For example:
//
//   x := ctx.Const(ctx.Symbol("x"), ctx.IntSort())
//   // ... further solving
//   m.Eval(x) => x's value
//
// Maps: Z3_model_eval
func (m *Model) Eval(c *AST) *AST {
	var result C.Z3_ast
	if C.Z3_model_eval(m.rawCtx, m.rawModel, c.rawAST, C.Z3_TRUE, &result) != C.Z3_TRUE {
		return nil
	}

	return &AST{
		rawCtx: m.rawCtx,
		rawAST: result,
	}
}

// Assignments returns a map of all the assignments for all the constants
// within the model. The key of the map will be the String value of the
// symbol.
//
// This doesn't map to any specific Z3 API. This is a higher-level function
// provided by go-z3 to make the Z3 API easier to consume in Go.
func (m *Model) Assignments() map[string]*AST {
	result := make(map[string]*AST)
	for i := uint(0); i < m.NumConsts(); i++ {
		// Get the declaration
		decl := m.ConstDecl(i)

		// Get the name of it, i.e. "x"
		name := decl.DeclName()

		// Get the assignment for this
		ast := C.Z3_model_get_const_interp(
			m.rawCtx, m.rawModel, C.Z3_to_func_decl(decl.rawCtx, decl.rawAST))

		// Map it
		result[name.String()] = &AST{
			rawCtx: m.rawCtx,
			rawAST: ast,
		}
	}

	return result
}

// NumConsts returns the number of constant assignments.
//
// Maps: Z3_model_get_num_consts
func (m *Model) NumConsts() uint {
	return uint(C.Z3_model_get_num_consts(m.rawCtx, m.rawModel))
}

// ConstDecl returns the const declaration for the given index. idx must
// be less than NumConsts.
//
// Maps: Z3_model_get_const_decl
func (m *Model) ConstDecl(idx uint) *AST {
	return &AST{
		rawCtx: m.rawCtx,
		rawAST: C.Z3_func_decl_to_ast(
			m.rawCtx,
			C.Z3_model_get_const_decl(m.rawCtx, m.rawModel, C.uint(idx))),
	}
}

//-------------------------------------------------------------------
// Memory Management
//-------------------------------------------------------------------

// Close decreases the reference count for this model. If nothing else
// has manually increased the reference count, this will free the memory
// associated with it.
func (m *Model) Close() error {
	C.Z3_model_dec_ref(m.rawCtx, m.rawModel)
	return nil
}

// IncRef increases the reference count of this model. This is advanced,
// you probably don't need to use this.
func (m *Model) IncRef() {
	C.Z3_model_inc_ref(m.rawCtx, m.rawModel)
}

// DecRef decreases the reference count of this model. This is advanced,
// you probably don't need to use this.
//
// Close will decrease it automatically from the initial 1, so this should
// only be called with exact matching calls to IncRef.
func (m *Model) DecRef() {
	C.Z3_model_dec_ref(m.rawCtx, m.rawModel)
}
