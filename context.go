package z3

// #include "go-z3.h"
import "C"

// Context is what handles most of the interactions with Z3.
type Context struct {
	raw C.Z3_context
}

func NewContext(c *Config) *Context {
	return &Context{
		raw: C.Z3_mk_context(c.Z3Value()),
	}
}

// Close frees the memory associated with this context.
func (c *Context) Close() error {
	// Clear context
	C.Z3_del_context(c.raw)

	// Clear error handling
	errorHandlerMapLock.Lock()
	delete(errorHandlerMap, c.raw)
	errorHandlerMapLock.Unlock()

	return nil
}

// Z3Value returns the internal structure for this Context.
func (c *Context) Z3Value() C.Z3_context {
	return c.raw
}
