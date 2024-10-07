package githubController

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/xorima/webhook-bridge/internal/data/topic"
	"io"
	"testing"
)

type mockProducer struct {
	event   *topic.Event
	channel *topic.Channel
	err     error
}

func (m *mockProducer) Produce(ctx context.Context, channel *topic.Channel, event *topic.Event) error {
	m.channel = channel
	m.event = event
	return m.err
}

func newMockProducer(err error) *mockProducer {
	return &mockProducer{err: err}
}

func NewMockBody(content string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(content)))
}

func TestNewMockBody(t *testing.T) {
	t.Run("it should return a body correctly", func(t *testing.T) {
		b := NewMockBody("hello-world")
		body, err := io.ReadAll(b)
		assert.NoError(t, err)
		assert.Equal(t, "hello-world", string(body))
	})
}

type FailingReadCloser struct{}

func (f *FailingReadCloser) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func (f *FailingReadCloser) Close() error {
	return nil
}

func TestMockProducer(t *testing.T) {
	t.Run("it should implement the producer interface", func(t *testing.T) {
		assert.Implements(t, (*topic.EventProducer)(nil), newMockProducer(nil))
	})
}
