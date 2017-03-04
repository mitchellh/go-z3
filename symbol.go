package z3

import (
	"strconv"
	"unsafe"
)

// #include <stdlib.h>
// #include "go-z3.h"
import "C"

// Symbol represents a named
type Symbol struct {
	rawCtx    C.Z3_context
	rawSymbol C.Z3_symbol
}

// Create a symbol named by a string within the context.
//
// The memory associated with this symbol is freed when the context is freed.
func (c *Context) Symbol(name string) *Symbol {
	ns := C.CString(name)
	defer C.free(unsafe.Pointer(ns))

	return &Symbol{
		rawCtx:    c.raw,
		rawSymbol: C.Z3_mk_string_symbol(c.raw, ns),
	}
}

// Create a symbol named by an int within the context.
//
// The memory associated with this symbol is freed when the context is freed.
func (c *Context) SymbolInt(name int) *Symbol {
	return &Symbol{
		rawCtx:    c.raw,
		rawSymbol: C.Z3_mk_int_symbol(c.raw, C.int(name)),
	}
}

// String returns a string value for this symbol no matter what kind
// of symbol it is. If it is an int, it will be converted to a string
// result.
func (s *Symbol) String() string {
	switch C.Z3_get_symbol_kind(s.rawCtx, s.rawSymbol) {
	case C.Z3_INT_SYMBOL:
		return strconv.FormatInt(
			int64(C.Z3_get_symbol_int(s.rawCtx, s.rawSymbol)), 10)

	case C.Z3_STRING_SYMBOL:
		// We don't need to free this value since it uses statically allocated
		// space that is reused by Z3. The GoString call will copy the memory.
		return C.GoString(C.Z3_get_symbol_string(s.rawCtx, s.rawSymbol))

	default:
		return "unknown symbol kind"
	}
}
