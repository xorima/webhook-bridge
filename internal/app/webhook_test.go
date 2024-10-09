package app

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
)

func TestWebhookHandler_Post(t *testing.T) {
	t.Run("it should return 400 when an error occours", func(t *testing.T) {
		handler := NewWebhookHandler(slogger.NewDevNullLogger(), &MockController{err: assert.AnError})
		req := httptest.NewRequest(http.MethodPost, "/v1/webhook/github", bytes.NewBuffer([]byte(`{}`)))
		rec := httptest.NewRecorder()
		handler.RegisterRoutes(chi.NewRouter(), nil)
		handler.Post(rec, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body, err := io.ReadAll(rec.Result().Body)
		assert.NoError(t, err)
		assert.JSONEq(t, `{"status": 400, "message": "bad request"}`, string(body))

	})
	t.Run("it should return 202 when there is not an error", func(t *testing.T) {
		handler := NewWebhookHandler(slogger.NewDevNullLogger(), &MockController{})
		req := httptest.NewRequest(http.MethodPost, "/v1/webhook/github", bytes.NewBuffer([]byte(`{}`)))
		rec := httptest.NewRecorder()
		handler.RegisterRoutes(chi.NewRouter(), nil)
		handler.Post(rec, req)
		assert.Equal(t, http.StatusAccepted, rec.Code)
		body, err := io.ReadAll(rec.Result().Body)
		assert.NoError(t, err)
		assert.JSONEq(t, `{"status": 202, "message": "accepted"}`, string(body))
	})
}
