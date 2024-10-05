package topic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvent(t *testing.T) {
	t.Run("it should create without issue", func(t *testing.T) {
		e := NewEvent("1.0.0", "hello-world")
		assert.Equal(t, "1.0.0", e.Version)
		assert.Equal(t, "hello-world", e.Body)
		assert.Empty(t, e.Attributes)
	})
	t.Run("it should add attributes without issue", func(t *testing.T) {
		e := NewEvent("1.0.0", "hello-world", NewAttribute("key", "value"))
		assert.Equal(t, "1.0.0", e.Version)
		assert.Equal(t, "hello-world", e.Body)
		assert.Len(t, e.Attributes, 1)
	})
}
