package redisClient

import (
	"context"
	"fmt"
	"github-bridge/internal/data/topic"
	"github.com/redis/go-redis/v9"
	"github.com/xorima/slogger"
	"log/slog"
	"strings"
)

const (
	versionKey    = "version"
	bodyKey       = "body"
	attrPrefixKey = "attr_"
)

func (c *Client) CreateConsumerGroup(ctx context.Context, channel *topic.Channel, consumer *topic.Consumer) error {
	return c.client.XGroupCreateMkStream(ctx, channelName(channel), consumer.Group, "0").Err()
}

func (c *Client) Consume(ctx context.Context, channel *topic.Channel, consumer *topic.Consumer) ([]*topic.Event, error) {
	var streams []string
	var events []*topic.Event
	streams = append(streams, channelName(channel), c.consumeType(consumer.MessageStrategy))
	c.log.Debug("streams configured", slog.Any("streams", streams))
	consumers := c.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    consumer.Group,
		Consumer: consumer.Name,
		Streams:  streams,
	})
	err := consumers.Err()
	if err != nil {
		return nil, err
	}
	for _, stream := range consumers.Val() {
		for _, message := range stream.Messages {
			events = append(events, c.messageToEvent(message))
		}
	}
	return events, nil
}

func (c *Client) Produce(ctx context.Context, channel *topic.Channel, event *topic.Event) error {
	err := c.client.XAdd(ctx, &redis.XAddArgs{
		Stream: channelName(channel),
		Values: eventBody(event),
	}).Err()
	if err != nil {
		c.log.Error("error encountered while writing event to channel", slog.String("channel", channelName(channel)), slogger.ErrorAttr(err))
		return err
	}
	return nil
}

func (c *Client) ClearChannel(ctx context.Context, channel *topic.Channel) error {
	return c.client.Del(ctx, channelName(channel)).Err()
}

func (c *Client) consumeType(strategy topic.ConsumerStrategy) string {
	// https://medium.com/redis-with-raphael-de-lio/understanding-redis-streams-33aa96ca7206
	switch strategy {
	case topic.ConsumerStrategyAllTime:
		c.log.Debug("consumer set to all time", slog.String("strategy-value", "0"))
		// The ID from which the group will start reading. If you want the group to read all elements since the beginning, use “0” as the ID
		return "0"
	default:
		c.log.Debug("consumer set to all time", slog.String("strategy-value", ">"))
		// > means that the consumer should read new elements that haven’t been read by any other consumer in the group.
		return ">"
	}
}

func channelName(channel *topic.Channel) string {
	var nameParts []string
	nameParts = append(nameParts, channel.Prefix...)
	nameParts = append(nameParts, channel.Name)
	return strings.Join(nameParts, ":")
}

func eventBody(event *topic.Event) map[string]any {
	response := map[string]any{
		versionKey: event.Version,
		bodyKey:    event.Body,
	}
	for _, a := range event.Attributes {
		response[fmt.Sprintf("%s%s", attrPrefixKey, a.Key)] = a.Value
	}
	return response
}

func (c *Client) messageToEvent(msg redis.XMessage) *topic.Event {
	var body, version string
	var attributes []topic.Attribute

	for key, v := range msg.Values {
		if key == bodyKey {
			body = v.(string)
			continue
		}
		if key == versionKey {
			version = v.(string)
			continue
		}
		if strings.Contains(key, attrPrefixKey) {
			attributes = append(attributes, topic.NewAttribute(strings.ReplaceAll(key, attrPrefixKey, ""), v))
			continue
		}
	}
	return topic.NewEvent(version, body, attributes...)
}
