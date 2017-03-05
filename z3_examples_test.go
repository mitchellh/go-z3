package z3

import (
	"fmt"
)

// From C examples: find_model_example2
func ExampleFindModel2() {
	// Create the context
	config := NewConfig()
	defer config.Close()
	ctx := NewContext(config)
	defer ctx.Close()

	// Create the solver
	s := ctx.NewSolver()
	defer s.Close()

	// Create a couple variables
	x := ctx.Const(ctx.Symbol("x"), ctx.IntSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.IntSort())

	// Create a couple integers
	v1 := ctx.Const(ctx.SymbolInt(1), ctx.IntSort())
	v2 := ctx.Const(ctx.SymbolInt(2), ctx.IntSort())

	// y + 1
	y_plus_one := y.Add(v1)

	// x < y + 1 && x > 2
	c1 := x.Lt(y_plus_one)
	c2 := x.Gt(v2)

	// Assert the constraints
	s.Assert(c1)
	s.Assert(c2)

	{
		// Solve
		fmt.Println("Solving part 1")
		if v := s.Check(); v != True {
			fmt.Println("unsatisfied!")
			return
		}

		// Get the resulting model:
		m := s.Model()
		fmt.Printf("Model:\n%s", m)
		m.Close()
	}

	// Create some new assertions
	//
	// !(x == y)
	c3 := x.Eq(y).Not()
	s.Assert(c3)

	{
		// Solve
		fmt.Println("\nSolving part 2")
		if v := s.Check(); v != True {
			fmt.Println("unsatisfied!")
			return
		}

		// Get the resulting model:
		m := s.Model()
		fmt.Printf("Model:\n%s", m)
		m.Close()
	}

	// Output:
	// Solving part 1
	// Model:
	// x -> 0
	// k!2 -> (- 1)
	// k!1 -> 0
	// y -> 1
	//
	// Solving part 2
	// Model:
	// k!1 -> 0
	// k!2 -> (- 1)
	// y -> 1
	// x -> 0
}
