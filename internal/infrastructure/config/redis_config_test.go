package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"os"
	"testing"
)

func TestAppConfig_RedisHostname(t *testing.T) {
	t.Run("it should return an empty string when nothing is set", func(t *testing.T) {
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "", cfg.RedisConfig().Hostname())
	})

	t.Run("it should return redis.example.com when env is set to redis.example.com", func(t *testing.T) {
		t.Setenv("REDIS_HOSTNAME", "redis.example.com")
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "redis.example.com", cfg.RedisConfig().Hostname())
	})

	t.Run("it should return from config file when set to redis.example.org", func(t *testing.T) {
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("redis:\n  hostname: redis.example.org"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "redis.example.org", cfg.RedisConfig().Hostname())
	})

	t.Run("it should take from env when both env and config file are set", func(t *testing.T) {
		t.Setenv("REDIS_HOSTNAME", "redis.example.com")
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("redis:\n  hostname: redis.example.org"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "redis.example.com", cfg.RedisConfig().Hostname())
	})
}

func TestAppConfig_RedisPassword(t *testing.T) {
	t.Run("it should return an empty string when nothing is set", func(t *testing.T) {
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "", cfg.RedisConfig().Password())
	})

	t.Run("it should return secret when env is set to secret", func(t *testing.T) {
		t.Setenv("REDIS_PASSWORD", "secret")
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "secret", cfg.RedisConfig().Password())
	})

	t.Run("it should return from config file when set to supersecret", func(t *testing.T) {
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("redis:\n  password: supersecret"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "supersecret", cfg.RedisConfig().Password())
	})

	t.Run("it should take from env when both env and config file are set", func(t *testing.T) {
		t.Setenv("REDIS_PASSWORD", "secret")
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("redis:\n  password: supersecret"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "secret", cfg.RedisConfig().Password())
	})
}

func TestAppConfig_RedisDB(t *testing.T) {
	t.Run("it should return 0 when nothing is set", func(t *testing.T) {
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, 0, cfg.RedisConfig().DB())
	})

	t.Run("it should return 1 when env is set to 1", func(t *testing.T) {
		t.Setenv("REDIS_DB", "1")
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, 1, cfg.RedisConfig().DB())
	})

	t.Run("it should return from config file when set to 2", func(t *testing.T) {
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("redis:\n  db: 2"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, 2, cfg.RedisConfig().DB())
	})

	t.Run("it should take from env when both env and config file are set", func(t *testing.T) {
		t.Setenv("REDIS_DB", "1")
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("redis:\n  db: 2"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, 1, cfg.RedisConfig().DB())
	})
	t.Run("it should return 0 (int default value) when set to an invalid value", func(t *testing.T) {
		t.Setenv("REDIS_DB", "foobar")
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, 0, cfg.RedisConfig().DB())
	})
}
