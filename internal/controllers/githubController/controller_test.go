package githubController

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"net/http"
	"testing"
)

func TestNewController(t *testing.T) {
	t.Run("it should create without issue", func(t *testing.T) {
		c := NewController(slogger.NewDevNullLogger(), newMockProducer(nil))
		assert.NotNil(t, c.producer)
	})
}

func TestController_Process(t *testing.T) {
	t.Run("it should return an error if the github event header is missing", func(t *testing.T) {
		c := NewController(slogger.NewDevNullLogger(), newMockProducer(nil))
		err := c.Process(context.Background(), http.Header{}, NewMockBody("hello world"))
		assert.ErrorIs(t, err, ErrMissingHeader)
	})
	t.Run("it should return an error if it cannot parse the body", func(t *testing.T) {
		c := NewController(slogger.NewDevNullLogger(), newMockProducer(nil))
		headers := http.Header{}
		headers.Add(githubEventHeader, "pull-request")
		err := c.Process(context.Background(), headers, &FailingReadCloser{})
		assert.ErrorIs(t, err, ErrCannotReadBody)
	})
	t.Run("it should return an error if the producer has an error", func(t *testing.T) {
		c := NewController(slogger.NewDevNullLogger(), newMockProducer(assert.AnError))
		headers := http.Header{}
		headers.Add(githubEventHeader, "pull-request")
		err := c.Process(context.Background(), headers, NewMockBody("hello world"))
		assert.ErrorIs(t, err, ErrFailedToPublish)
	})
	t.Run("it should publish the body to the correct queue", func(t *testing.T) {
		p := newMockProducer(nil)
		c := NewController(slogger.NewDevNullLogger(), p)
		headers := http.Header{}
		headers.Add(githubEventHeader, "pull-request")
		err := c.Process(context.Background(), headers, NewMockBody("hello world"))
		assert.NoError(t, err)
		assert.Equal(t, "hello world", p.event.Body)
		assert.Equal(t, "pull-request", p.event.Attributes[0].Value)
		assert.Equal(t, "github-events", p.channel.Name)
	})
}
