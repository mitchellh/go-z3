package z3

import (
	"testing"
)

func TestModelAssignments(t *testing.T) {
	config := NewConfig()
	defer config.Close()

	ctx := NewContext(config)
	defer ctx.Close()

	// Create a symbol
	x := ctx.Const(ctx.Symbol("x"), ctx.IntSort())

	// x + 4 = 16
	ast := x.Add(ctx.Int(4, ctx.IntSort())).
		Eq(ctx.Int(16, ctx.IntSort()))
	t.Logf("\nAST:\n%s", ast.String())

	// Create the solver
	s := ctx.NewSolver()
	defer s.Close()

	// Assert constraints
	s.Assert(ast)

	// Solve
	result := s.Check()
	if result != True {
		t.Fatalf("bad: %s", result)
	}

	// Get the model
	m := s.Model()
	defer m.Close()
	t.Logf("\nModel:\n%s", m.String())

	// Get the exact value
	am := m.Assignments()
	assign := am["x"]
	t.Logf("Assignment: %s", assign)
	if assign.Int() != 12 {
		t.Fatalf("bad: %s", assign)
	}
}

func TestModelEval(t *testing.T) {
	config := NewConfig()
	defer config.Close()

	ctx := NewContext(config)
	defer ctx.Close()

	// Create a symbol
	x := ctx.Const(ctx.Symbol("x"), ctx.IntSort())

	// x + 4 = 16
	ast := x.Add(ctx.Int(4, ctx.IntSort())).
		Eq(ctx.Int(16, ctx.IntSort()))
	t.Logf("\nAST:\n%s", ast.String())

	// Create the solver
	s := ctx.NewSolver()
	defer s.Close()

	// Assert constraints
	s.Assert(ast)

	// Solve
	result := s.Check()
	if result != True {
		t.Fatalf("bad: %s", result)
	}

	// Get the model
	m := s.Model()
	defer m.Close()
	t.Logf("\nModel:\n%s", m.String())

	// Get the exact value
	assign := m.Eval(x)
	t.Logf("Assignment: %s", assign)
	if assign.Int() != 12 {
		t.Fatalf("bad: %s", assign)
	}
}
