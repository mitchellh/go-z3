package z3

import (
	"testing"
)

func TestASTDistinct(t *testing.T) {
	config := NewConfig()
	defer config.Close()
	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int
	x := ctx.Const(ctx.Symbol("x"), ctx.IntSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.IntSort())
	z := ctx.Const(ctx.Symbol("z"), ctx.IntSort())

	// Add
	raw := x.Distinct(y, z)

	actual := raw.String()
	if actual != "(distinct x y z)" {
		t.Fatalf("bad:\n%s", actual)
	}
}

func TestASTNot(t *testing.T) {
	config := NewConfig()
	defer config.Close()
	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int
	x := ctx.Const(ctx.Symbol("x"), ctx.BoolSort())

	// Add
	raw := x.Not()

	actual := raw.String()
	if actual != "(not x)" {
		t.Fatalf("bad:\n%s", actual)
	}
}

func TestASTEq(t *testing.T) {
	config := NewConfig()
	defer config.Close()
	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int
	x := ctx.Const(ctx.Symbol("x"), ctx.BoolSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.BoolSort())

	// Add
	raw := x.Eq(y)

	actual := raw.String()
	if actual != "(= x y)" {
		t.Fatalf("bad:\n%s", actual)
	}
}

func TestASTIte(t *testing.T) {
	config := NewConfig()
	defer config.Close()
	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int
	x := ctx.Const(ctx.Symbol("x"), ctx.BoolSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.BoolSort())
	z := ctx.Const(ctx.Symbol("z"), ctx.BoolSort())

	raw := x.Ite(y, z)

	actual := raw.String()
	if actual != "(ite x y z)" {
		t.Fatalf("bad:\n%s", actual)
	}
}

func TestASTIff(t *testing.T) {
	config := NewConfig()
	defer config.Close()
	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int
	x := ctx.Const(ctx.Symbol("x"), ctx.BoolSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.BoolSort())

	raw := x.Iff(y)

	actual := raw.String()
	if actual != "(= x y)" {
		t.Fatalf("bad:\n%s", actual)
	}
}

func TestASTImplies(t *testing.T) {
	config := NewConfig()
	defer config.Close()
	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int
	x := ctx.Const(ctx.Symbol("x"), ctx.BoolSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.BoolSort())

	// Add
	raw := x.Implies(y)

	actual := raw.String()
	if actual != "(=> x y)" {
		t.Fatalf("bad:\n%s", actual)
	}
}

func TestASTXor(t *testing.T) {
	config := NewConfig()
	defer config.Close()
	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int
	x := ctx.Const(ctx.Symbol("x"), ctx.BoolSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.BoolSort())

	// Add
	raw := x.Xor(y)

	actual := raw.String()
	if actual != "(xor x y)" {
		t.Fatalf("bad:\n%s", actual)
	}
}

func TestASTAnd(t *testing.T) {
	config := NewConfig()
	defer config.Close()
	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int
	x := ctx.Const(ctx.Symbol("x"), ctx.BoolSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.BoolSort())

	// Add
	raw := x.And(y)

	actual := raw.String()
	if actual != "(and x y)" {
		t.Fatalf("bad:\n%s", actual)
	}
}

func TestASTOr(t *testing.T) {
	config := NewConfig()
	defer config.Close()
	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int
	x := ctx.Const(ctx.Symbol("x"), ctx.BoolSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.BoolSort())

	// Add
	raw := x.Or(y)

	actual := raw.String()
	if actual != "(or x y)" {
		t.Fatalf("bad:\n%s", actual)
	}
}
