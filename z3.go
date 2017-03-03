package z3

// #cgo CFLAGS: -Ivendor/z3/src/api
// #cgo LDFLAGS: ${SRCDIR}/libz3.a -lstdc++
// #include <stdlib.h>
// #include "go-z3.h"
import "C"
