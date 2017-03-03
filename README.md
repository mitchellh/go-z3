# Go Bindings to the Z3 Theorem Prover

go-z3 provides bindings to [Z3](https://github.com/Z3Prover/z3), a
theorem prover out of Microsoft Research. Z3 is a state-of-the-art
[SMT Solver](https://en.wikipedia.org/wiki/Satisfiability_modulo_theories).

This library provides bindings via cgo directly to Z3.

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

## Issues and Contributing

If you find an issue with this library, please report an issue. If you'd like,
we welcome any contributions. Fork this library and submit a pull request.
