package main

import (
	"errors"
	"fmt"

	pkgerrors "github.com/pkg/errors"
)

func fn4() error {
	e1 := errors.New("error")

	// e1 := pkgerrors.New("error")
	return fn3(e1)
}

func fn1(e error) error {
	return pkgerrors.Wrap(e, "inner")
}

func fn2(e error) error {
	return fn1(pkgerrors.Wrap(e, "middle"))
}

func fn3(e error) error {
	return fn2(pkgerrors.Wrap(e, "outer"))
}

func main() {
	type stackTracer interface {
		StackTrace() pkgerrors.StackTrace
	}

	err, ok := pkgerrors.Cause(fn4()).(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := err.StackTrace()
	fmt.Printf("%+v", st[0:2]) // top two frames

	// Example output:
	// github.com/pkg/errors_test.fn
	//	/home/dfc/src/github.com/pkg/errors/example_test.go:47
	// github.com/pkg/errors_test.Example_stackTrace
	//	/home/dfc/src/github.com/pkg/errors/example_test.go:127
}
