// Package z3 provides Go bindings to the Z3 SMT Solver.
//
// The bindings are a balance between idiomatic Go and being an obvious
// translation from the Z3 API so you can look up Z3 APIs and find the
// intuitive mapping in Go.
//
// The most foreign thing to Go programmers will be error handling. Rather
// than return the `error` type from almost every function, the z3 package
// mimics Z3's API by requiring you to set an error handler callback. This
// error handler will be invoked whenenver an error occurs. See
// ErrorHandler and Context.SetErrorHandler for more information.
package z3

// #cgo CFLAGS: -Ivendor/z3/src/api
// #cgo LDFLAGS: ${SRCDIR}/libz3.a -lstdc++ -lm
// #include <stdlib.h>
// #include "go-z3.h"
import "C"
