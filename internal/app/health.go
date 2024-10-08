package app

import (
	"log/slog"
	"net/http"
)

type HealthHandler struct {
	log *slog.Logger
}

func NewHealthHandler(log *slog.Logger) *HealthHandler {
	return &HealthHandler{log: log.With(slog.String("handler", "health"))}
}

// Get - The app Health
//
//	@Summary		This API is there to receive the health of this instance.
//	@Description	Returns the health of this instance.
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	app.Response	"Healthy"
//	@Failure		500	{object}	app.Response	"Internal Server Error"
//	@Router			/healthz [get]
func (h *HealthHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.log.InfoContext(r.Context(), "returning health check", slog.String("user-agent", r.UserAgent()))
	NewResponse(http.StatusOK, "Healthy").WriteResponse(w)
}

func (h *HealthHandler) RegisterRoutes(r Router) {
	r.Get("/healthz", h.Get)
}
