package z3

import (
	"unsafe"
)

// #include <stdlib.h>
// #include "go-z3.h"
import "C"

// Config is used to set configuration for Z3. This should be created with
// NewConfig and closed with Close when you're done using it.
//
// Config structures are used to set parameters for Z3 behavior. See the
// Z3 docs for information on available parameters. They can be set with
// SetParamValue.
//
// As for 2016-03-02, the parameters available are documented as:
//
// proof (Boolean) Enable proof generation
// debug_ref_count (Boolean) Enable debug support for Z3_ast reference counting
// trace (Boolean) Tracing support for VCC
// trace_file_name (String) Trace out file for VCC traces
// timeout (unsigned) default timeout (in milliseconds) used for solvers
// well_sorted_check type checker
// auto_config use heuristics to automatically select solver and configure it
// model model generation for solvers, this parameter can be overwritten when creating a solver
// model_validate validate models produced by solvers
// unsat_core unsat-core generation for solvers, this parameter can be overwritten when creating a solver
//
type Config struct {
	raw C.Z3_config
}

// NewConfig allocates a new configuration object.
func NewConfig() *Config {
	return &Config{
		raw: C.Z3_mk_config(),
	}
}

// Close frees the memory associated with this configuration
func (c *Config) Close() error {
	C.Z3_del_config(c.raw)
	return nil
}

// SetParamValue sets the parameters for a Config. See the Config docs.
func (c *Config) SetParamValue(k, v string) {
	ck := C.CString(k)
	cv := C.CString(v)

	// We free the strings since they're not actually stored
	defer C.free(unsafe.Pointer(ck))
	defer C.free(unsafe.Pointer(cv))

	C.Z3_set_param_value(c.raw, ck, cv)
}

// Z3Value returns the raw internal pointer value. This should only be
// used if you really understand what you're doing. It may be invalid after
// Close is called.
func (c *Config) Z3Value() C.Z3_config {
	return c.raw
}
