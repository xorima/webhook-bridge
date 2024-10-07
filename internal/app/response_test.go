package app

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestNewResponse(t *testing.T) {
	t.Run("it should create the response correctly and lowercase the message", func(t *testing.T) {
		r := NewResponse(200, "Hello world")
		assert.Equal(t, 200, r.Status)
		assert.Equal(t, "hello world", r.Message)
	})
}

func TestResponse_ToJson(t *testing.T) {
	t.Run("it should serialize to json correctly", func(t *testing.T) {
		r := NewResponse(400, "bad request")
		expectedJSON := `{"status":400,"message":"bad request"}`
		assert.JSONEq(t, expectedJSON, string(r.ToJson()))
	})
}

func TestResponse_WriteResponse(t *testing.T) {
	t.Run("it should write the response correctly", func(t *testing.T) {
		r := NewResponse(400, "bad request")
		rr := httptest.NewRecorder()
		r.WriteResponse(rr)
		assert.Equal(t, 400, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		assert.Equal(t, "nosniff", rr.Header().Get("X-Content-Type-Options"))
		expectedJSON := `{"status":400,"message":"bad request"}`
		assert.JSONEq(t, expectedJSON, rr.Body.String())
	})
}
