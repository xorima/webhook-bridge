package app

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"github.com/xorima/webhook-bridge/internal/infrastructure/config"
)

type MockController struct {
	err error
}

func (m *MockController) Process(ctx context.Context, headers http.Header, body io.ReadCloser) error {
	return m.err
}

func TestApp_Run(t *testing.T) {
	logger := slogger.NewDevNullLogger()
	mockController := &MockController{}
	mockConfig := &config.AppConfig{}

	app := NewApp(logger, mockController, mockConfig)

	// Channel to signal when the server is ready
	ready := make(chan struct{})

	// Run the server in a separate goroutine
	go func() {
		err := app.Run()
		assert.NoError(t, err)
	}()

	// Wait for the server to start
	go func() {
		for {
			resp, err := http.Get("http://localhost:3000")
			if err == nil {
				resp.Body.Close()
				close(ready)
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Wait for the ready signal
	<-ready

	// Check if the server is running by making a request
	resp, err := http.Get("http://localhost:3000")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Send a signal to stop the server
	p, err := os.FindProcess(os.Getpid())
	assert.NoError(t, err)
	err = p.Signal(os.Interrupt)
	assert.NoError(t, err)

	// Loop to check if the server has stopped
	stopped := false
	for i := 0; i < 10; i++ {
		_, err = http.Get("http://localhost:3000")
		if err != nil {
			stopped = true
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	assert.True(t, stopped, "Server did not stop in the expected time")
	// Check if the server has stopped by making a request
	_, err = http.Get("http://localhost:3000")
	assert.Error(t, err)
}
