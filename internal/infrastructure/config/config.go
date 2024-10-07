package config

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/xorima/slogger"
	"log/slog"
)

type AppCfg interface {
	LogLevel() string
	Version() string
	Hostname() string
	RedisConfig() RedisCfg
	GitHubConfig() GitHubCfg
}

const (
	logLevelKey    = "log.level"
	apiVersionKey  = "api.version"
	apiHostnameKey = "api.hostname"
)

type AppConfig struct {
	v            *viper.Viper
	redisConfig  RedisCfg
	githubConfig GitHubCfg
}

func setupViper(v *viper.Viper, cfgFile string) {
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath("/etc/app/")
		v.AddConfigPath(".")
	}
}

func NewAppConfig(log *slog.Logger, cfgFile string) (*AppConfig, error) {
	log = slogger.SubLogger(log, "config")
	v := viper.New()
	setupViper(v, cfgFile)
	v.AutomaticEnv()

	bindEnvAndDefault(v, logLevelKey, "info")
	bindEnvAndDefault(v, apiVersionKey, "dev")
	bindEnvAndDefault(v, apiHostnameKey, "localhost:3000")

	err := v.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Warn("unable to load from config file, continuing with env vars", slogger.ErrorAttr(err))
		} else {
			log.Error("a critical error has been found during config loading", slogger.ErrorAttr(err))
			return nil, err
		}
	}
	return &AppConfig{v: v,
		redisConfig:  newRedisConfig(v),
		githubConfig: newGitHubConfig(v),
	}, nil
}

func (c *AppConfig) LogLevel() string {
	return c.v.GetString(logLevelKey)
}

func (c *AppConfig) Version() string {
	return c.v.GetString(apiVersionKey)
}

func (c *AppConfig) Hostname() string {
	return c.v.GetString(apiHostnameKey)
}

func (c *AppConfig) RedisConfig() RedisCfg {
	return c.redisConfig
}

func (c *AppConfig) GitHubConfig() GitHubCfg {
	return c.githubConfig
}
