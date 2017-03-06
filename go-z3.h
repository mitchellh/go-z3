// vim: ft=c ts=2 sts=2 st=2
/*
 * This header exists to simplify the headers that are included within
 * the Go files. This header should include all the necessary headers
 * for the compilation of the Go library.
 * */

#ifndef _GOZ3_H_INCLUDED
#define _GOZ3_H_INCLUDED

#include <z3.h>

//-------------------------------------------------------------------
// Error handling helpers
//-------------------------------------------------------------------
// This is declared in error.go and is a way for us to call back into
// Go to execute the proper error handlers.
extern void goZ3ErrorHandler(Z3_context, Z3_error_code);

// This method is used as a way to get a valid error handler
// pointer back into Go.
static inline Z3_error_handler* _go_z3_error_handler() {
    return &goZ3ErrorHandler;
}

#endif
