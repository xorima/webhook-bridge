package config

import "github.com/spf13/viper"

type RedisCfg interface {
	Hostname() string
	Password() string
	DB() int
}

const (
	redisHostnameKey = "redis.hostname"
	redisPasswordKey = "redis.password"
	redisDbKey       = "redis.db"
)

type RedisConfig struct {
	v *viper.Viper
}

func newRedisConfig(v *viper.Viper) *RedisConfig {
	bindEnvAndDefault(v, redisHostnameKey, nil)
	bindEnvAndDefault(v, redisPasswordKey, nil)
	bindEnvAndDefault(v, redisDbKey, 0)
	return &RedisConfig{v: v}
}

func (c *RedisConfig) Hostname() string {
	return c.v.GetString(redisHostnameKey)
}

func (c *RedisConfig) Password() string {
	return c.v.GetString(redisPasswordKey)
}

func (c *RedisConfig) DB() int {
	return c.v.GetInt(redisDbKey)
}
