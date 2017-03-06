package z3

import "sync"

// #include "go-z3.h"
import "C"

// ErrorCode represents the enum of error codes Z3 supports.
type ErrorCode uint

const (
	ErrorCodeOk              ErrorCode = C.Z3_OK
	ErrorCodeSortError                 = C.Z3_SORT_ERROR
	ErrorCodeIOB                       = C.Z3_IOB
	ErrorCodeInvalidArg                = C.Z3_INVALID_ARG
	ErrorCodeParserError               = C.Z3_PARSER_ERROR
	ErrorCodeNoParser                  = C.Z3_NO_PARSER
	ErrorCodeInvalidPattern            = C.Z3_INVALID_PATTERN
	ErrorCodeMemoutFail                = C.Z3_MEMOUT_FAIL
	ErrorCodeFileAccessError           = C.Z3_FILE_ACCESS_ERROR
	ErrorCodeInternalFatal             = C.Z3_INTERNAL_FATAL
	ErrorCodeInvalidUsage              = C.Z3_INVALID_USAGE
	ErrorCodeDecRefError               = C.Z3_DEC_REF_ERROR
	ErrorCodeException                 = C.Z3_EXCEPTION
)

// ErrorHandler is the callback that is invoked when an error occurs in
// Z3 and is registered by SetErrorHandler.
type ErrorHandler func(*Context, ErrorCode)

// These unexported vars are used to keep track of our error handlers.
var errorHandlerMap = map[C.Z3_context]ErrorHandler{}
var errorHandlerMapLock sync.RWMutex

// SetErrorHandler registers the error handler. This handler is invoked
// whenever an error occurs within Z3.
func (c *Context) SetErrorHandler(f ErrorHandler) {
	C.Z3_set_error_handler(c.raw, C._go_z3_error_handler())

	errorHandlerMapLock.Lock()
	defer errorHandlerMapLock.Unlock()
	errorHandlerMap[c.raw] = f
}

// Error returns the error message for the given error code.
// This code can be retrieved via the error handler callback.
//
// This MUST be called during the handler. This must not be called later
// since the error state on the context may have cleared.
//
// Maps: Z3_get_error_msg_ex
func (c *Context) Error(code ErrorCode) string {
	return C.GoString(C.Z3_get_error_msg_ex(c.raw, C.Z3_error_code(code)))
}

//export goZ3ErrorHandler
func goZ3ErrorHandler(raw C.Z3_context, code C.Z3_error_code) {
	errorHandlerMapLock.RLock()
	defer errorHandlerMapLock.RUnlock()

	// Look up the error handler for this context
	f, ok := errorHandlerMap[raw]
	if !ok {
		return
	}

	// Call it!
	f(&Context{raw: raw}, ErrorCode(code))
}
