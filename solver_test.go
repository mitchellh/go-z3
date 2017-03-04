package z3

import (
	"testing"
)

func TestSolver(t *testing.T) {
	config := NewConfig()
	defer config.Close()

	ctx := NewContext(config)
	defer ctx.Close()

	// Create the "x xor y" constraint
	boolTyp := ctx.BoolSort()
	x := ctx.Const(ctx.Symbol("x"), boolTyp)
	y := ctx.Const(ctx.Symbol("y"), boolTyp)
	x_xor_y := x.Xor(y)
	ast := x_xor_y
	t.Logf("\nAST:\n%s", ast.String())

	// Create the solver
	s := ctx.NewSolver()
	defer s.Close()

	// Assert constraints
	s.Assert(x_xor_y)

	// Solve
	result := s.Check()
	if result != True {
		t.Fatalf("bad: %s", result)
	}

	// Get the model
	m := s.Model()
	defer m.Close()
	t.Logf("\nModel:\n%s", m.String())
}
