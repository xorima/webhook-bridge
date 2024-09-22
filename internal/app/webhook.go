package app

import (
	"log/slog"
	"net/http"
)

type WebhookHandler struct {
	log *slog.Logger
}

func NewWebhookHandler(log *slog.Logger) *WebhookHandler {
	return &WebhookHandler{
		log: log.With(slog.String("handler", "webhook")),
	}
}

// Post - Receives the GitHub webhooks
//
//	@Summary		This API is there to receive the GitHub events.
//	@Description	Once received, this will add them to the event stream for consumers.
//	@Tags			Webhooks
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	app.Response	"Successful Response"
//	@Failure		401	{object}	app.Response		"Unauthorized"
//	@Failure		404	{object}	app.Response		"Failure Response"
//	@Router			/api/v1/webhook/github [post]
func (wh *WebhookHandler) Post(w http.ResponseWriter, r *http.Request) {
	wh.log.Info("got request")
	resp := NewResponse(http.StatusAccepted, "Accepted")
	w.WriteHeader(resp.Status)
	w.Write(resp.ToJson())
}

func (wh *WebhookHandler) RegisterRoutes(r Router) {
	r.Post("/v1/webhook/github", wh.Post)
}
