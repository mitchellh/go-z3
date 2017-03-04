package z3

// #include "go-z3.h"
import "C"

// LBool is the lifted boolean type representing false, true, and undefined.
type LBool int8

const (
	False LBool = C.Z3_L_FALSE
	Undef       = C.Z3_L_UNDEF
	True        = C.Z3_L_TRUE
)
