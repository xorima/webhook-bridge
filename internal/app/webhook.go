package app

import (
	"github.com/xorima/slogger"
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
//	@Success		202	{object} Response	"Successful Response"
//	@Failure		    401	{object} Response	"Unauthorized"
//	@Failure		  404	{object} Response		"Failure Response"
//	@Router			/api/v1/webhook/github [post]
func (wh *WebhookHandler) Post(w http.ResponseWriter, r *http.Request) {
	wh.log.Info("got request")
	resp := NewResponse(http.StatusAccepted, "Accepted")
	w.WriteHeader(resp.Status)
	_, err := w.Write(resp.ToJson())
	if err != nil {
		wh.log.Error("failure in writing webhook response", slogger.ErrorAttr(err))
	}
}

func (wh *WebhookHandler) RegisterRoutes(r Router) {
	r.Post("/v1/webhook/github", wh.Post)
}
