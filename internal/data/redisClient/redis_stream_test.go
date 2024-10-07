package redisClient

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"github.com/xorima/webhook-bridge/internal/data/topic"
	"github.com/xorima/webhook-bridge/internal/testhelpers/containers"
	"testing"
)

func TestClient_Produce(t *testing.T) {
	t.Parallel()
	t.Run("it should write the body to the correct stream/channel and be able to read it", func(t *testing.T) {
		ctx := context.Background()
		cont, err := containers.NewRedisContainer(ctx)
		assert.NoError(t, err)
		assert.NoError(t, cont.Start(ctx))
		defer func(cont *containers.RedisContainer, ctx context.Context) {
			err := cont.Stop(ctx)
			assert.NoError(t, err)
		}(cont, ctx)
		c := NewClient(newMockRedisCfg(fmt.Sprintf("127.0.0.1:%d", cont.GetRedisPort(ctx)), 0), slogger.NewDevNullLogger())
		attrs := []topic.Attribute{topic.NewAttribute("type", "foo"), topic.NewAttribute("logging-key", "123456789")}
		version, body := "1.0.0", "hello-world"
		event := topic.NewEvent(version, body, attrs...)
		channel := topic.NewChannel("testclient").WithPrefix("this", "is", "a", "test")
		consumer := topic.NewConsumer("foo", "bar")

		err = c.Produce(ctx, channel, event)
		assert.NoError(t, err)
		assert.NoError(t, c.CreateConsumerGroup(ctx, channel, consumer))
		events, err := c.Consume(ctx, channel, consumer)
		assert.NoError(t, err)

		assert.Len(t, events, 1)
		assert.Equal(t, version, events[0].Version)
		assert.Equal(t, body, events[0].Body)
		assert.Len(t, events[0].Attributes, 2)
		attrMap := make(map[string]any)
		for _, attr := range events[0].Attributes {
			attrMap[attr.Key] = attr.Value
		}
		assert.Equal(t, "foo", attrMap["type"])
		assert.Equal(t, "123456789", attrMap["logging-key"])
		assert.NoError(t, c.ClearChannel(context.Background(), channel))
	})
	t.Run("it should error when it cannot consume a message", func(t *testing.T) {
		ctx := context.Background()
		cont, err := containers.NewRedisContainer(ctx)
		assert.NoError(t, err)
		assert.NoError(t, cont.Start(ctx))
		defer func(cont *containers.RedisContainer, ctx context.Context) {
			err := cont.Stop(ctx)
			assert.NoError(t, err)
		}(cont, ctx)
		c := NewClient(newMockRedisCfg(fmt.Sprintf("127.0.0.1:%d", cont.GetRedisPort(ctx)), 0), slogger.NewDevNullLogger())
		attrs := []topic.Attribute{topic.NewAttribute("type", "foo"), topic.NewAttribute("logging-key", "123456789")}
		version, body := "1.0.0", "hello-world"
		event := topic.NewEvent(version, body, attrs...)
		channel := topic.NewChannel("testclient").WithPrefix("this", "is", "a", "test")
		// close the connection so it will error
		assert.NoError(t, c.Close())
		err = c.Produce(ctx, channel, event)
		assert.Error(t, err)
	})
}

func TestClient_Consume(t *testing.T) {
	t.Run("it should error if it cannot connect to the channel to consume", func(t *testing.T) {
		ctx := context.Background()
		cont, err := containers.NewRedisContainer(ctx)
		assert.NoError(t, err)
		assert.NoError(t, cont.Start(ctx))
		defer func(cont *containers.RedisContainer, ctx context.Context) {
			err := cont.Stop(ctx)
			assert.NoError(t, err)
		}(cont, ctx)
		c := NewClient(newMockRedisCfg(fmt.Sprintf("127.0.0.1:%d", cont.GetRedisPort(ctx)), 0), slogger.NewDevNullLogger())
		channel := topic.NewChannel("testclient").WithPrefix("this", "is", "a", "test")
		consumer := topic.NewConsumer("foo", "bar")
		_, err = c.Consume(ctx, channel, consumer)
		assert.Error(t, err) // channel not created so cannot consume
	})
	t.Run("it should consume all messages since last connection", func(t *testing.T) {
		ctx := context.Background()
		cont, err := containers.NewRedisContainer(ctx)
		assert.NoError(t, err)
		assert.NoError(t, cont.Start(ctx))
		defer func(cont *containers.RedisContainer, ctx context.Context) {
			err := cont.Stop(ctx)
			assert.NoError(t, err)
		}(cont, ctx)
		c := NewClient(newMockRedisCfg(fmt.Sprintf("127.0.0.1:%d", cont.GetRedisPort(ctx)), 0), slogger.NewDevNullLogger())
		attrs := []topic.Attribute{topic.NewAttribute("type", "foo"), topic.NewAttribute("logging-key", "123456789")}
		version, body := "1.0.0", "hello-world"
		event := topic.NewEvent(version, body, attrs...)
		channel := topic.NewChannel("testclient").WithPrefix("this", "is", "a", "test")
		consumer := topic.NewConsumer("foo", "bar")

		err = c.Produce(ctx, channel, event)
		assert.NoError(t, err)
		assert.NoError(t, c.CreateConsumerGroup(ctx, channel, consumer))
		events, err := c.Consume(ctx, channel, consumer)
		assert.NoError(t, err)

		assert.Len(t, events, 1)
		assert.Equal(t, version, events[0].Version)
		assert.Equal(t, body, events[0].Body)
		assert.Len(t, events[0].Attributes, 2)
		attrMap := make(map[string]any)
		for _, attr := range events[0].Attributes {
			attrMap[attr.Key] = attr.Value
		}
		assert.Equal(t, "foo", attrMap["type"])
		assert.Equal(t, "123456789", attrMap["logging-key"])
		assert.NoError(t, c.ClearChannel(context.Background(), channel))
	})
	// there seems to be a condition where the individual consumer in the group must have connected with new only BEFORE the 0 will work.
	t.Run("it should consume all messages possible", func(t *testing.T) {
		ctx := context.Background()
		cont, err := containers.NewRedisContainer(ctx)
		assert.NoError(t, err)
		assert.NoError(t, cont.Start(ctx))
		defer func(cont *containers.RedisContainer, ctx context.Context) {
			err := cont.Stop(ctx)
			assert.NoError(t, err)
		}(cont, ctx)
		c := NewClient(newMockRedisCfg(fmt.Sprintf("127.0.0.1:%d", cont.GetRedisPort(ctx)), 0), slogger.NewDevNullLogger())
		attrs := []topic.Attribute{topic.NewAttribute("type", "foo"), topic.NewAttribute("logging-key", "123456789")}
		version, body := "1.0.0", "hello-world"
		event := topic.NewEvent(version, body, attrs...)
		channel := topic.NewChannel("testclient").WithPrefix("this", "is", "a", "test")
		consumer := topic.NewConsumer("foo", "bar")
		assert.NoError(t, c.CreateConsumerGroup(ctx, channel, consumer))
		assert.NoError(t, c.Produce(ctx, channel, event))
		assert.NoError(t, c.Produce(ctx, channel, event))

		events, err := c.Consume(ctx, channel, consumer)
		assert.NoError(t, err)
		assert.Len(t, events, 2)
		// this will re-get the two above that have been delivered before
		events, err = c.Consume(ctx, channel, consumer.WithConsumeNewFromAllTime())
		assert.NoError(t, err)
		assert.Len(t, events, 2)
	})
}
