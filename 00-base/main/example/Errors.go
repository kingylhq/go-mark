package example

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and
// have type `error`, a built-in interface.
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` constructs a basic `error` value
		// with the given error message.
		return -1, errors.New("can't work with 42")
	}

	// A `nil` value in the error position indicates that
	// there was no error.
	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is used to
// signify a specific error condition.
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// We can wrap errors with higher-level errors to add
		// context. The simplest way to do this is with the
		// `%w` verb in `fmt.Errorf`. Wrapped errors
		// create a logical chain (A wraps B, which wraps C, etc.)
		// that can be queried with functions like `errors.Is`
		// and `errors.As`.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func Errors() {
	for _, i := range []int{7, 42} {

		// It's common to use an inline error check in the `if`
		// line.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	// 定义一个int数组
	arr := []int{1, 2, 3, 4, 5}
	for i := range arr {
		if err := makeTea(i); err != nil {

			// `errors.Is` checks that a given error (or any error in its chain)
			// matches a specific error value. This is especially useful with wrapped or
			// nested errors, allowing you to identify specific error types or sentinel
			// errors in a chain of errors.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
