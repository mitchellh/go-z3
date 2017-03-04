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

// String returns a human-friendly string version of the model.
func (m *Model) String() string {
	return C.GoString(C.Z3_model_to_string(m.rawCtx, m.rawModel))
}
