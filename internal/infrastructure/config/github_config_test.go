package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"os"
	"testing"
)

func TestAppConfig_GitHubHmacEnabled(t *testing.T) {
	t.Run("it should return true when nothing is set", func(t *testing.T) {
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.True(t, cfg.GitHubConfig().HmacEnabled())
	})

	t.Run("it should return false when env is set to false", func(t *testing.T) {
		t.Setenv("GITHUB_HMAC_ENABLED", "false")
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.False(t, cfg.GitHubConfig().HmacEnabled())
	})

	t.Run("it should return from config file when set to false", func(t *testing.T) {
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("github:\n  hmac:\n    enabled: false"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.False(t, cfg.GitHubConfig().HmacEnabled())
	})

	t.Run("it should take from env when both env and config file are set", func(t *testing.T) {
		t.Setenv("GITHUB_HMAC_ENABLED", "true")
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("github:\n  hmac:\n    enabled: false"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.True(t, cfg.GitHubConfig().HmacEnabled())
	})
}

func TestAppConfig_GitHubHmacSecret(t *testing.T) {
	t.Run("it should return an empty string when nothing is set", func(t *testing.T) {
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "", cfg.GitHubConfig().HmacSecret())
	})

	t.Run("it should return secret when env is set to secret", func(t *testing.T) {
		t.Setenv("GITHUB_HMAC_SECRET", "secret")
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "secret", cfg.GitHubConfig().HmacSecret())
	})

	t.Run("it should return from config file when set to supersecret", func(t *testing.T) {
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("github:\n  hmac:\n    secret: supersecret"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "supersecret", cfg.GitHubConfig().HmacSecret())
	})

	t.Run("it should take from env when both env and config file are set", func(t *testing.T) {
		t.Setenv("GITHUB_HMAC_SECRET", "secret")
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("github:\n  hmac:\n    secret: supersecret"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "secret", cfg.GitHubConfig().HmacSecret())
	})
}
