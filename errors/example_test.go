package errors_test

import (
	"errors"
	"fmt"

	xerrors "github.com/gechr/x/errors"
)

func ExampleIsAny() {
	errNotFound := errors.New("not found")
	errUnavailable := errors.New("unavailable")
	err := fmt.Errorf("lookup: %w", errNotFound)

	fmt.Println(xerrors.IsAny(err, errNotFound, errUnavailable))
	// Output: true
}
