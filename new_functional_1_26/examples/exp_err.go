package examples

import (
	"errors"
	"fmt"
)

type MyError struct {
	Msg string
}

func (e *MyError) Error() string { return e.Msg }

func ErrType() {
	var myErr *MyError

	err := error(&MyError{Msg: "boom"})

	handle := func(e *MyError) {
		fmt.Println("handled:", e.Error())
	}

	// Before Go 1.26
	if errors.As(err, &myErr) {
		handle(myErr)
	}

	// Go 1.26
	if myErr, ok := errors.AsType[*MyError](err); ok {
		handle(myErr)
	}

}
