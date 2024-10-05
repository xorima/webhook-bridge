package topic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAttribute(t *testing.T) {
	t.Run("it should create without issue", func(t *testing.T) {
		a := NewAttribute("key", "value")
		assert.Equal(t, "key", a.Key)
		assert.Equal(t, "value", a.Value)
	})
}
