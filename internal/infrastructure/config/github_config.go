package config

import "github.com/spf13/viper"

type GitHubCfg interface {
	HmacEnabled() bool
	HmacSecret() string
}

const (
	githubHmacEnabledKey = "github.hmac.enabled"
	githubHmacSecretKey  = "github.hmac.secret"
)

type GitHubConfig struct {
	v *viper.Viper
}

func newGitHubConfig(v *viper.Viper) *GitHubConfig {
	bindEnvAndDefault(v, githubHmacEnabledKey, true)
	bindEnvAndDefault(v, githubHmacSecretKey, nil)
	return &GitHubConfig{v: v}
}

func (c *GitHubConfig) HmacEnabled() bool {
	return c.v.GetBool(githubHmacEnabledKey)
}

func (c *GitHubConfig) HmacSecret() string {
	return c.v.GetString(githubHmacSecretKey)
}
