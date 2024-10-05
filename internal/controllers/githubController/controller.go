package githubController

import (
	"github.com/xorima/slogger"
	"log/slog"
)

type Controller struct {
	log *slog.Logger
}

func NewController(log *slog.Logger) *Controller {
	return &Controller{log: slogger.SubLogger(log, "githubController")}
}

// check the type based on header, put to right queue.
