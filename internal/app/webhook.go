package app

import "net/http"

type WebhookHandler struct {
}

type Webhook struct {
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{}
}

// Post - Receives the GitHub webhooks
//
//	@Summary		This API is there to receive the GitHub events.
//	@Description	Once received, this will add them to the event stream for consumers.
//	@Tags			GitHub webhook
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	app.Webhook	"Successful Response"
//	@Failure		401	{object}	app.Webhook		"Unauthorized"
//	@Failure		404	{object}	app.Webhook		"Failure Response"
//	@Router			/api/v1/webhook/github [post]
func (wh *WebhookHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("event created"))
}

func (wh *WebhookHandler) RegisterRoutes(r Router) {
	r.Post("/v1/webhook/github", wh.Post)
	r.Get("/v1/webhook/github", wh.Post)
}
