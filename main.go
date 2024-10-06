package main

import (
	"github.com/xorima/slogger"
	"github.com/xorima/webhook-bridge/docs"
	"github.com/xorima/webhook-bridge/internal/app"
	"os"
)

// @title           Webhook Bridge API
// @description     This is a bridge to receive various webhook events and publish them to a channel.

// @contact.name   Jason Field
// @contact.url    https://github.com/xorima

// @license.name  MIT
// @license.url  https://github.com/xorima/webhook-bridge/blob/main/LICENSE

// @host      localhost:3000
// @BasePath  /

// @externalDocs.description  GitHub
// @externalDocs.url          https://github.com/xorima/webhook-bridge
func main() {
	loggerOpts := slogger.NewLoggerOpts(
		"github.com/xorima/webhook-bridge",
		"github.com/xorima/webhook-bridge")
	logger := slogger.NewLogger(loggerOpts, slogger.WithLevel("debug"))
	logger.Info("starting app")
	h := app.NewApp(logger)
	docs.SwaggerInfo.Version = getVersion()
	docs.SwaggerInfo.Host = getHost()
	err := h.Run()
	if err != nil {
		logger.Error("runtime error", slogger.ErrorAttr(err))
	}
}

func getVersion() string {
	v := os.Getenv("VERSION")
	if len(v) > 1 {
		return v
	}
	return "dev"
}

func getHost() string {
	v := os.Getenv("API_HOST")
	if len(v) > 1 {
		return v
	}
	return "localhost:3000"
}
