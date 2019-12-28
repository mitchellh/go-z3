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

func (b LBool) String() string {
	if b == False {
		return "False"
	} else if b == Undef {
		return "Undef"
	} else if b == True {
		return "True"
	} else {
		panic("Unknown LBool")
	}
}
