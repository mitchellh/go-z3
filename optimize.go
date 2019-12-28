package z3

// #include "go-z3.h"
import "C"

type Optimize struct {
	rawCtx      C.Z3_context
	rawOptimize C.Z3_optimize
}

// NewOptimize creates a new optimize.
func (c *Context) NewOptimize() *Optimize {
	rawOptimize := C.Z3_mk_optimize(c.raw)
	C.Z3_optimize_inc_ref(c.raw, rawOptimize)

	return &Optimize{
		rawOptimize: rawOptimize,
		rawCtx:      c.raw,
	}
}

// Close frees the memory associated with this.
func (s *Optimize) Close() error {
	C.Z3_optimize_dec_ref(s.rawCtx, s.rawOptimize)
	return nil
}

// Add adds a constraint onto the Optimize.
//
// Maps to: Z3_optimize_assert
func (s *Optimize) Add(a *AST) {
	C.Z3_optimize_assert(s.rawCtx, s.rawOptimize, a.rawAST)
}

// Maximize adds a function to maximize and returns a handle to be later used in Lower or Upper.
//
// Maps to: Z3_optimize_maximize
func (s *Optimize) Maximize(a *AST) (handle uint) {
	return uint(C.Z3_optimize_maximize(s.rawCtx, s.rawOptimize, a.rawAST))
}

// Minimize adds a function to minimize and returns a handle to be later used in Lower or Upper.
//
// Maps to: Z3_optimize_minimize
func (s *Optimize) Minimize(a *AST) (handle uint) {
	return uint(C.Z3_optimize_minimize(s.rawCtx, s.rawOptimize, a.rawAST))
}

// Lower returns a lower value or the current approximation.
//
// Maps to: Z3_optimize_get_lower
func (s *Optimize) Lower(handle uint) *AST {
	return &AST{
		rawCtx: s.rawCtx,
		rawAST: C.Z3_optimize_get_lower(s.rawCtx, s.rawOptimize, C.uint(handle)),
	}
}

// Upper returns an upper value or the current approximation.
//
// Maps to: Z3_optimize_get_upper
func (s *Optimize) Upper(handle uint) *AST {
	return &AST{
		rawCtx: s.rawCtx,
		rawAST: C.Z3_optimize_get_upper(s.rawCtx, s.rawOptimize, C.uint(handle)),
	}
}

// Check checks if the currently set formula is consistent.
//
// Maps to: Z3_optimize_check
func (s *Optimize) Check() LBool {
	return LBool(C.Z3_optimize_check(s.rawCtx, s.rawOptimize))
}

// Model returns the last model from a Check.
//
// Maps to: Z3_optimize_get_model
func (s *Optimize) Model() *Model {
	m := &Model{
		rawCtx:   s.rawCtx,
		rawModel: C.Z3_optimize_get_model(s.rawCtx, s.rawOptimize),
	}
	m.IncRef()
	return m
}
