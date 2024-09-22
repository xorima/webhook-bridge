package main

import (
	"github-bridge/internal/app"
	"github.com/xorima/slogger"
)

func main() {
	loggerOpts := slogger.NewLoggerOpts(
		"github-bridge",
		"github-bridge")
	logger := slogger.NewLogger(loggerOpts, slogger.WithLevel("debug"))
	logger.Info("starting app")
	h := app.NewApp(logger)
	err := h.Run()
	if err != nil {
		logger.Error("runtime error", slogger.ErrorAttr(err))
	}
}
