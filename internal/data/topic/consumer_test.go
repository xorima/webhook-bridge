package topic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	t.Run("it should create with strategy of since last group connection by default", func(t *testing.T) {
		c := NewConsumer("name", "group")
		assert.Equal(t, "name", c.Name)
		assert.Equal(t, "group", c.Group)
		assert.Equal(t, ConsumerStrategyUndelivered, c.MessageStrategy)
	})
}

func TestConsumer_WithConsumeNewFromAllTime(t *testing.T) {
	t.Run("it should set it to the all time strategy", func(t *testing.T) {
		c := NewConsumer("name", "group").WithConsumeNewFromAllTime()
		assert.Equal(t, "name", c.Name)
		assert.Equal(t, "group", c.Group)
		assert.Equal(t, ConsumerStrategyAllTime, c.MessageStrategy)
	})
}
