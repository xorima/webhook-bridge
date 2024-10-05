package redisClient

import (
	"github.com/redis/go-redis/v9"
	"github.com/xorima/slogger"
	"log/slog"
)

type Client struct {
	client *redis.Client
	log    *slog.Logger
}

func NewClient(addr, password string, db int, log *slog.Logger) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &Client{
		client: rdb,
		log:    slogger.SubLogger(log, "redis"),
	}
}

func (c *Client) Close() error {
	err := c.client.Close()
	if err != nil {
		c.log.Error("Failed to close Redis client", slogger.ErrorAttr(err))
	}
	return err
}
