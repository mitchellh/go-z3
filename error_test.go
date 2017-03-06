package z3

import (
	"strings"
	"testing"
)

func TestContextErrorHandler(t *testing.T) {
	config := NewConfig()
	defer config.Close()
	ctx := NewContext(config)
	defer ctx.Close()

	// Set an error handler
	called := false
	msg := ""
	ctx.SetErrorHandler(func(c *Context, code ErrorCode) {
		called = true
		msg = c.Error(code)
	})

	// Create an int
	x := ctx.Const(ctx.Symbol("x"), ctx.BoolSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.BoolSort())

	// This won't work because x and y aren't ints
	x.Ge(y)
	if !called {
		t.Fatal("should call error handler")
	}
	if !strings.Contains(msg, "Sort mismatch") {
		t.Fatalf("bad: %s", msg)
	}
}
