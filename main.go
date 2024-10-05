package main

import (
	"github.com/xorima/slogger"
	"github.com/xorima/webhook-bridge/internal/app"
)

func main() {
	loggerOpts := slogger.NewLoggerOpts(
		"github.com/xorima/webhook-bridge",
		"github.com/xorima/webhook-bridge")
	logger := slogger.NewLogger(loggerOpts, slogger.WithLevel("debug"))
	logger.Info("starting app")
	h := app.NewApp(logger)
	err := h.Run()
	if err != nil {
		logger.Error("runtime error", slogger.ErrorAttr(err))
	}
}
