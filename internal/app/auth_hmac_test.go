package app

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockHmacConfig struct {
	secret  string
	enabled bool
}

func (m *MockHmacConfig) HmacSecret() string {
	return m.secret
}

func (m *MockHmacConfig) HmacEnabled() bool {
	return m.enabled
}

type mockReader struct {
}

func (r *mockReader) Read(p []byte) (n int, err error) {
	return 0, assert.AnError
}

func TestAuthHmacMiddleware(t *testing.T) {
	t.Run("it should not check hmac when disabled", func(t *testing.T) {
		mockCfg := &MockHmacConfig{
			secret:  "test-secret",
			enabled: false,
		}
		authHmac := NewAuthHmacMiddleware(slogger.NewDevNullLogger(), mockCfg)
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"key":"value"}`))
		rr := httptest.NewRecorder()
		handler := authHmac.AuthHmacMiddleware(nextHandler)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("it should return bad request when body is invalid/cannot be read", func(t *testing.T) {
		mockCfg := &MockHmacConfig{
			secret:  "test-secret",
			enabled: true,
		}
		authHmac := NewAuthHmacMiddleware(slogger.NewDevNullLogger(), mockCfg)

		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodPost, "/", &mockReader{})
		rr := httptest.NewRecorder()

		handler := authHmac.AuthHmacMiddleware(nextHandler)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("it should return 401 when hmac is not valid", func(t *testing.T) {
		mockCfg := &MockHmacConfig{
			secret:  "test-secret",
			enabled: true,
		}
		authHmac := NewAuthHmacMiddleware(slogger.NewDevNullLogger(), mockCfg)

		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"key":"value"}`))
		req.Header.Set("X-Hub-Signature-256", "invalid-hmac")
		rr := httptest.NewRecorder()

		handler := authHmac.AuthHmacMiddleware(nextHandler)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("it should continue to serve when hmac is valid", func(t *testing.T) {
		mockCfg := &MockHmacConfig{
			secret:  "test-secret",
			enabled: true,
		}
		authHmac := NewAuthHmacMiddleware(slogger.NewDevNullLogger(), mockCfg)

		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		body := []byte("hello-world")

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		// https://www.freeformatter.com/hmac-generator.html#before-output
		req.Header.Set("X-Hub-Signature-256", "sha256=08de25c4a512c01793d703e75503b69a954ecaa1c2fc8917237f8f1e5d5ddd7c")
		rr := httptest.NewRecorder()

		handler := authHmac.AuthHmacMiddleware(nextHandler)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}
