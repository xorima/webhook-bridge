package config

import (
	"github.com/spf13/viper"
	"strings"
)

func bindEnvAndDefault(v *viper.Viper, key string, d any) {
	if d != nil {
		v.SetDefault(key, d)
	}
	_ = v.BindEnv(key, strings.ReplaceAll(strings.ToUpper(key), ".", "_"))
}
