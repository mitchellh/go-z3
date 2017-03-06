# Go Bindings to the Z3 Theorem Prover

go-z3 provides bindings to [Z3](https://github.com/Z3Prover/z3), a
theorem prover out of Microsoft Research. Z3 is a state-of-the-art
[SMT Solver](https://en.wikipedia.org/wiki/Satisfiability_modulo_theories).

This library provides bindings via cgo directly to Z3.

**Library Status:** Most major sections of the API are covered, but the
Z3 API is massive. Adding missing APIs should be trivial unless the
entire category/type is not implemented yet. Please issue a pull request
for any missing APIs and I'll add it!

## Installation

Installation is a little trickier than a standard Go library, but not
by much. You can't simply `go get` this library, unfortunately. This is
because [Z3](https://github.com/Z3Prover/z3) must first be built. We
don't ship a pre-built version of Z3.

To build Z3, we've made it very easy. You will need the following packages
available on your host operating system:

* Python
* Make
* gcc/Clang

Then just type:

```
$ make
```

This will download Z3, compile it, and run the tests for go-z3,
verifying that your build is functional. By default, go-z3 will download
and build the "master" ref of Z3, but this is customizable.

Compiling/installing the go-z3 library should work on Linux, Mac OS X,
and Windows. On Windows, msys is the only supported build toolchain (same
as Go itself).

**Due to this linking, it is strongly recommended that you vendor this
repository and bake our build system into your process.**

### Customizing the Z3 Compilation

You can customize the Z3 compilation by setting a couple environmental
variables prior to calling `make`:

  * `Z3_REF` is the git ref that will be checked out for Z3. This
    defaults to to a recently tagged version. It is recommended that you
    explicitly set this to a ref that works for you to avoid any changes
    in this library later.

## Usage

go-z3 exposes the Z3 API in a style that mostly idiomatic Go. The API
should be comfortable to use by any Go programmer without having intimate
knowledge of how Z3 works.

For usage examples and documentation, please see the
[go-z3 GoDoc](http://godoc.org/github.com/mitchellh/go-z3), which
we keep up to date and full of examples.

For a quick taste of what using go-z3 looks like, though, we provide
a basic example below:

```go
package main

import (
	"fmt"

	"github.com/mitchellh/go-z3"
)

func main() {
	// Create the context
	config := z3.NewConfig()
	ctx := z3.NewContext(config)
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
```

## Issues and Contributing

If you find an issue with this library, please report an issue. If you'd like,
we welcome any contributions. Fork this library and submit a pull request.
