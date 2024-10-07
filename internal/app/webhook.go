package app

import (
	"github.com/xorima/slogger"
	"github.com/xorima/webhook-bridge/internal/controllers"
	"log/slog"
	"net/http"
)

type WebhookHandler struct {
	log  *slog.Logger
	ctrl controllers.Controller
}

func NewWebhookHandler(log *slog.Logger, controller controllers.Controller) *WebhookHandler {
	return &WebhookHandler{
		log:  log.With(slog.String("handler", "webhook")),
		ctrl: controller,
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
//	@Failure		  400	{object} Response		"Bad Request"
//	@Router			/api/v1/webhook/github [post]
func (wh *WebhookHandler) Post(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse(http.StatusAccepted, "Accepted")
	wh.log.InfoContext(r.Context(), "got request")
	err := wh.ctrl.Process(r.Context(), r.Header, r.Body)
	if err != nil {
		wh.log.ErrorContext(r.Context(), "failure during processing of event", slogger.ErrorAttr(err))
		resp = NewResponse(http.StatusBadRequest, "Bad Request")
	}
	w.WriteHeader(resp.Status)
	_, err = w.Write(resp.ToJson())
	if err != nil {
		wh.log.ErrorContext(r.Context(), "failure in writing webhook response", slogger.ErrorAttr(err))
	}
}

func (wh *WebhookHandler) RegisterRoutes(r Router) {
	r.Post("/v1/webhook/github", wh.Post)
}
