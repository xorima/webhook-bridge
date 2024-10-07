package redisClient

import (
	"github.com/redis/go-redis/v9"
	"github.com/xorima/slogger"
	"github.com/xorima/webhook-bridge/internal/infrastructure/config"
	"log/slog"
)

type Client struct {
	client *redis.Client
	log    *slog.Logger
}

func NewClient(cfg config.RedisCfg, log *slog.Logger) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Hostname(),
		Password: cfg.Password(),
		DB:       cfg.DB(),
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
