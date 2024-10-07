package errs

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapError(t *testing.T) {
	t.Run("it should wrap the error correctly", func(t *testing.T) {
		e := WrapError(errors.New("hello world"), errors.New("new context for this error"))
		assert.ErrorContains(t, e, "hello world")
		assert.ErrorContains(t, e, "new context for this error")
	})
	t.Run("it should be understandable from IsError", func(t *testing.T) {
		err := fmt.Errorf("hello world")
		e := WrapError(err, errors.New("new context for this error"))
		assert.ErrorContains(t, e, "hello world")
		assert.ErrorContains(t, e, "new context for this error")
		assert.True(t, IsError(e, err))
	})
}
