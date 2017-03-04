package z3

// #include "go-z3.h"
import "C"

// Solver is a single solver tied to a specific Context within Z3.
//
// It is created via the NewSolver methods on Context. When a solver is
// no longer needed, the Close method must be called. This will remove the
// solver from the context and no more APIs on Solver may be called
// thereafter.
//
// Freeing the context (Context.Close) will NOT automatically close associated
// solvers. They must be managed separately.
type Solver struct {
	rawCtx    C.Z3_context
	rawSolver C.Z3_solver
}

// NewSolver creates a new solver.
func (c *Context) NewSolver() *Solver {
	rawSolver := C.Z3_mk_solver(c.raw)
	C.Z3_solver_inc_ref(c.raw, rawSolver)

	return &Solver{
		rawSolver: rawSolver,
		rawCtx:    c.raw,
	}
}

// Close frees the memory associated with this.
func (s *Solver) Close() error {
	C.Z3_solver_dec_ref(s.rawCtx, s.rawSolver)
	return nil
}

// Assert asserts a constraint onto the Solver.
//
// Maps to: Z3_solver_assert
func (s *Solver) Assert(a *AST) {
	C.Z3_solver_assert(s.rawCtx, s.rawSolver, a.rawAST)
}

// Check checks if the currently set formula is consistent.
//
// Maps to: Z3_solver_check
func (s *Solver) Check() LBool {
	return LBool(C.Z3_solver_check(s.rawCtx, s.rawSolver))
}

// Model returns the last model from a Check.
//
// Maps to: Z3_solver_get_model
func (s *Solver) Model() *Model {
	m := &Model{
		rawCtx:   s.rawCtx,
		rawModel: C.Z3_solver_get_model(s.rawCtx, s.rawSolver),
	}
	m.IncRef()
	return m
}
