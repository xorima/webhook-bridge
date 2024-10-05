package topic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChannel(t *testing.T) {
	t.Run("it should create without prefix", func(t *testing.T) {
		c := NewChannel("foo")
		assert.Equal(t, "foo", c.Name)
		assert.Empty(t, c.Prefix)
	})
}

func TestChannel_WithPrefix(t *testing.T) {
	t.Run("it should add the prefix", func(t *testing.T) {
		c := NewChannel("foo")
		assert.Equal(t, "foo", c.Name)
		assert.Empty(t, c.Prefix)
		c.WithPrefix("hello", "world")
		assert.Len(t, c.Prefix, 2)
		assert.Equal(t, "hello", c.Prefix[0])
		assert.Equal(t, "world", c.Prefix[1])
	})
}
