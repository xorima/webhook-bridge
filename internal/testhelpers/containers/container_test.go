package containers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContainer(t *testing.T) {
	t.Run("it will create the new container without error", func(t *testing.T) {
		ctx := context.Background()
		c, err := NewContainer(ctx, "redis:latest", "Ready to accept connections", []string{"6379/tcp"})
		assert.NoError(t, err)
		assert.NoError(t, c.Start(ctx))
		port, err := c.GetPort(ctx, "6379/tcp")
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, port, 1)
		assert.NoError(t, c.Stop(ctx))
	})
	t.Run("it should throw an error if the image is not supplied", func(t *testing.T) {
		t.Run("it throws an error if the container cannot be started", func(t *testing.T) {
			_, err := NewContainer(context.Background(), "", "", nil)
			assert.ErrorContains(t, err, "you must specify either a build context or an image")
		})
	})
}

func TestContainer_GetPort(t *testing.T) {
	t.Run("it should throw an error if the port was not defined at creation", func(t *testing.T) {
		ctx := context.Background()
		c, err := NewContainer(ctx, "redis:latest", "Ready to accept connections", []string{"6379/tcp"})
		assert.NoError(t, err)
		_, err = c.GetPort(ctx, "tcp/1234")
		assert.ErrorContains(t, err, "not found in defined ports at runtime")
	})
}
