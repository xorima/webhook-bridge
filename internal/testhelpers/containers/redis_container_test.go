package containers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRedisContainer(t *testing.T) {
	t.Run("it will create the redis container without error", func(t *testing.T) {
		ctx := context.Background()
		c, err := NewRedisContainer(ctx)
		assert.NoError(t, err)
		assert.NoError(t, c.Start(ctx))
		port := c.GetRedisPort(ctx)
		assert.GreaterOrEqual(t, port, 1)
		assert.NoError(t, c.Stop(ctx))
	})
}
