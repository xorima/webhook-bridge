package config

import (
	"github.com/xorima/slogger"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppConfig_Version(t *testing.T) {
	t.Run("it should return dev when nothing is set", func(t *testing.T) {
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "dev", cfg.Version())
	})

	t.Run("it should return 1.0.0 when env is set to 1.0.0", func(t *testing.T) {
		t.Setenv("API_VERSION", "1.0.0")
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "1.0.0", cfg.Version())
	})

	t.Run("it should return from config file when set to 1.2.3", func(t *testing.T) {
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("api:\n  version: 1.2.3"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "1.2.3", cfg.Version())
	})

	t.Run("it should take from env when both env and config file are set", func(t *testing.T) {
		t.Setenv("API_VERSION", "1.0.0")
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("api:\n  version: 1.2.3"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "1.0.0", cfg.Version())
	})
	t.Run("it should throw an error if the config exists but is not formatted correctly", func(t *testing.T) {
		t.Setenv("API_VERSION", "1.0.0")
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("this is not valid"), 0644)
		assert.NoError(t, err)

		_, err = NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.ErrorContains(t, err, "unmarshal errors")

	})

}
func TestAppConfig_Hostname(t *testing.T) {
	t.Run("it should return localhost:3000 when nothing is set", func(t *testing.T) {
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "localhost:3000", cfg.Hostname())
	})

	t.Run("it should return example.com when env is set to example.com", func(t *testing.T) {
		t.Setenv("API_HOSTNAME", "example.com")
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "example.com", cfg.Hostname())
	})

	t.Run("it should return from config file when set to example.org", func(t *testing.T) {
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("api:\n  hostname: example.org"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "example.org", cfg.Hostname())
	})

	t.Run("it should take from env when both env and config file are set", func(t *testing.T) {
		t.Setenv("API_HOSTNAME", "example.com")
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("hostname: example.org"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "example.com", cfg.Hostname())
	})
	t.Run("it should throw an error if the config exists but is not formatted correctly", func(t *testing.T) {
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("this is not valid"), 0644)
		assert.NoError(t, err)

		_, err = NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.ErrorContains(t, err, "unmarshal errors")
	})
}

func TestAppConfig_LogLevel(t *testing.T) {
	t.Run("it should return info when nothing is set", func(t *testing.T) {
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "info", cfg.LogLevel())
	})

	t.Run("it should return debug when env is set to debug", func(t *testing.T) {
		t.Setenv("LOG_LEVEL", "debug")
		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), "")
		assert.NoError(t, err)
		assert.Equal(t, "debug", cfg.LogLevel())
	})

	t.Run("it should return from config file when set to warn", func(t *testing.T) {
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("log:\n  level: warn"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "warn", cfg.LogLevel())
	})

	t.Run("it should take from env when both env and config file are set", func(t *testing.T) {
		t.Setenv("LOG_LEVEL", "debug")
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("log:\n  level: warn"), 0644)
		assert.NoError(t, err)

		cfg, err := NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "debug", cfg.LogLevel())
	})
	t.Run("it should throw an error if the config exists but is not formatted correctly", func(t *testing.T) {
		t.Setenv("LOG_LEVEL", "debug")
		tmpDir := t.TempDir()
		tmpFile := tmpDir + "/config.yaml"
		err := os.WriteFile(tmpFile, []byte("this is not valid"), 0644)
		assert.NoError(t, err)

		_, err = NewAppConfig(slogger.NewDevNullLogger(), tmpFile)
		assert.ErrorContains(t, err, "unmarshal errors")
	})
}
