package z3

import (
	"fmt"
)

// This example is a basic mathematical example
func ExampleBasicMath() {
	// Create the context
	config := NewConfig()
	ctx := NewContext(config)
	config.Close()
	defer ctx.Close()

	// Logic:
	// x + y + z > 4
	// x + y < 2
	// z > 0
	// x != y != z
	// x, y, z != 0
	// x + y = -3

	// Create the solver
	s := ctx.NewSolver()
	defer s.Close()

	// Vars
	x := ctx.Const(ctx.Symbol("x"), ctx.IntSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.IntSort())
	z := ctx.Const(ctx.Symbol("z"), ctx.IntSort())

	zero := ctx.Int(0, ctx.IntSort()) // To save repeats

	// x + y + z > 4
	s.Assert(x.Add(y, z).Gt(ctx.Int(4, ctx.IntSort())))

	// x + y < 2
	s.Assert(x.Add(y).Lt(ctx.Int(2, ctx.IntSort())))

	// z > 0
	s.Assert(z.Gt(zero))

	// x != y != z
	s.Assert(x.Distinct(y, z))

	// x, y, z != 0
	s.Assert(x.Eq(zero).Not())
	s.Assert(y.Eq(zero).Not())
	s.Assert(z.Eq(zero).Not())

	// x + y = -3
	s.Assert(x.Add(y).Eq(ctx.Int(-3, ctx.IntSort())))

	if v := s.Check(); v != True {
		fmt.Println("Unsolveable")
		return
	}

	// Get the resulting model:
	m := s.Model()
	assignments := m.Assignments()
	m.Close()
	fmt.Printf("x = %s\n", assignments["x"])
	fmt.Printf("y = %s\n", assignments["y"])
	fmt.Printf("z = %s\n", assignments["z"])

	// Output:
	// x = (- 2)
	// y = (- 1)
	// z = 8
}

// From C examples: demorgan
func ExampleDemorgan() {
	// Create the context
	config := NewConfig()
	ctx := NewContext(config)
	config.Close()
	defer ctx.Close()

	// Create a couple variables
	x := ctx.Const(ctx.Symbol("x"), ctx.BoolSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.BoolSort())

	// Final goal: !(x && y) == (!x || !y)
	// Built incrementally so its clearer

	// !(x && y)
	not_x_and_y := x.And(y).Not()

	// (!x || !y)
	not_x_or_not_y := x.Not().Or(y.Not())

	// Conjecture and negated
	conj := not_x_and_y.Iff(not_x_or_not_y)
	negConj := conj.Not()

	// Create the solver
	s := ctx.NewSolver()
	defer s.Close()

	// Assert the constraints
	s.Assert(negConj)

	if v := s.Check(); v == False {
		fmt.Println("DeMorgan is valid")
		return
	}

	// Output:
	// DeMorgan is valid
}

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
	v1 := ctx.Int(1, ctx.IntSort())
	v2 := ctx.Int(2, ctx.IntSort())

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
		assignments := m.Assignments()
		m.Close()
		fmt.Printf("x = %s\n", assignments["x"])
		fmt.Printf("y = %s\n", assignments["y"])
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
		assignments := m.Assignments()
		m.Close()
		fmt.Printf("x = %s\n", assignments["x"])
		fmt.Printf("y = %s", assignments["y"])
	}

	// Output:
	// Solving part 1
	// x = 3
	// y = 3
	//
	// Solving part 2
	// x = 3
	// y = 4
}
