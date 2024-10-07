package errs

import (
	"errors"
	"fmt"
)

func WrapError(originalErr, wrappedErr error) error {
	return fmt.Errorf("%w: %w", wrappedErr, originalErr)
}

func IsError(err, isError error) bool {
	return errors.Is(err, isError)
}
