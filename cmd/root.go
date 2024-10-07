package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xorima/slogger"
	"github.com/xorima/webhook-bridge/docs"
	"github.com/xorima/webhook-bridge/internal/app"
	"github.com/xorima/webhook-bridge/internal/controllers/githubController"
	"github.com/xorima/webhook-bridge/internal/data/redisClient"
	"github.com/xorima/webhook-bridge/internal/infrastructure/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "webhook-bridge",
	Short: "A webserver that bridges webhooks to event producers",
	RunE: func(cmd *cobra.Command, args []string) error {
		loggerOpts := slogger.NewLoggerOpts(
			"webhook-bridge",
			"webhook-bridge")
		cfg, err := config.NewAppConfig(slogger.NewLogger(loggerOpts), cfgFile)
		if err != nil {
			return err
		}
		logger := slogger.NewLogger(loggerOpts, slogger.WithLevel(cfg.LogLevel()))
		logger.Info("starting app")
		h := app.NewApp(logger,
			githubController.NewController(logger,
				redisClient.NewClient(cfg.RedisConfig(), logger),
				"local", "webhook", "bridge",
			),
		)
		docs.SwaggerInfo.Version = cfg.Version()
		docs.SwaggerInfo.Host = cfg.Hostname()
		err = h.Run()
		if err != nil {
			logger.Error("runtime error", slogger.ErrorAttr(err))
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/app/config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.AutomaticEnv() // read in environment variables that match
}
