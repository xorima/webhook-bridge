package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler_RegisterRoutes(t *testing.T) {
	router := chi.NewRouter()
	healthHandler := NewHealthHandler(slogger.NewDevNullLogger())
	healthHandler.RegisterRoutes(router)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"status":200,"message":"healthy"}`, rr.Body.String())
}
