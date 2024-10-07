package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/xorima/slogger"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwaggerHandler_Redirect(t *testing.T) {
	t.Run("it should redirect the  given path the swagger ui", func(t *testing.T) {
		swaggerHandler := NewSwaggerHandler(slogger.NewDevNullLogger())

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()
		swaggerHandler.Redirect(rr, req)
		assert.Equal(t, http.StatusFound, rr.Code)
		assert.Equal(t, "/swagger/", rr.Header().Get("Location"))
	})
}
func TestSwaggerHandler_RegisterRoutes(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{
			name: "it should redirect / to the swagger ui",
			path: "/",
		},
		{
			name: "it should redirect /swagger to the swagger ui",
			path: "/swagger",
		},
		{
			name: "it should redirect /swagger-ui to the swagger ui",
			path: "/swagger-ui",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := chi.NewRouter()
			swaggerHandler := NewSwaggerHandler(slogger.NewDevNullLogger())
			swaggerHandler.RegisterRoutes(router)

			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusFound, rr.Code)
			assert.Equal(t, "/swagger/", rr.Header().Get("Location"))
		})
	}
	t.Run("it should load the swagger ui", func(t *testing.T) {
		router := chi.NewRouter()
		swaggerHandler := NewSwaggerHandler(slogger.NewDevNullLogger())
		swaggerHandler.RegisterRoutes(router)
		req := httptest.NewRequest(http.MethodGet, "/swagger/", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusMovedPermanently, rr.Code)
		// this is a path mounted directly on the swagger ui that we do not control
		assert.Equal(t, "/swagger/index.html", rr.Header().Get("Location"))
	})

}
