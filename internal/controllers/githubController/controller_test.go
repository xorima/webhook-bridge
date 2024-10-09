package githubController

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"github.com/xorima/webhook-bridge/internal/controllers"
	"github.com/xorima/webhook-bridge/internal/data/topic"
	"net/http"
	"testing"
)

func TestNewController(t *testing.T) {
	t.Run("it should create without issue", func(t *testing.T) {
		c := NewController(slogger.NewDevNullLogger(), newMockProducer(nil))
		assert.NotNil(t, c.producer)
	})
	t.Run("it should implement the controller interface", func(t *testing.T) {
		assert.Implements(t, (*controllers.Controller)(nil), NewController(slogger.NewDevNullLogger(), newMockProducer(nil)))
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
		headers.Add(githubEventHeader, "PullRequest")
		err := c.Process(context.Background(), headers, &FailingReadCloser{})
		assert.ErrorIs(t, err, ErrCannotReadBody)
	})
	t.Run("it should return an error if the producer has an error", func(t *testing.T) {
		c := NewController(slogger.NewDevNullLogger(), newMockProducer(assert.AnError))
		headers := http.Header{}
		headers.Add(githubEventHeader, "PullRequest")
		headers.Add(githubDeliveryHeader, "1234")
		err := c.Process(context.Background(), headers, NewMockBody(`{"action":"labelled"}`))
		assert.ErrorIs(t, err, ErrFailedToPublish)
	})
	t.Run("it should continue without error if the action key is not there", func(t *testing.T) {
		p := newMockProducer(nil)
		c := NewController(slogger.NewDevNullLogger(), p, "test", "webhook", "bridge")
		body := `{"repo":"xorima/test-action"}`
		headers := http.Header{}
		headers.Add(githubEventHeader, "PullRequest")
		headers.Add(githubDeliveryHeader, "1234")
		err := c.Process(context.Background(), headers, NewMockBody(body))
		assert.NoError(t, err)
		assert.Equal(t, body, p.event.Body)
		assert.Equal(t, "pull-request", p.channel.Name)
		assert.Equal(t, []string{"test", "webhook", "bridge", "github"}, p.channel.Prefix)
		assert.Len(t, p.event.Attributes, 1)
		assert.Contains(t, p.event.Attributes, topic.NewAttribute("delivery-id", "1234"))
	})
	t.Run("it should continue without error if the delivery id is not found", func(t *testing.T) {
		p := newMockProducer(nil)
		c := NewController(slogger.NewDevNullLogger(), p, "test", "webhook", "bridge")
		body := `{"repo":"xorima/test-action"}`
		headers := http.Header{}
		headers.Add(githubEventHeader, "PullRequest")
		err := c.Process(context.Background(), headers, NewMockBody(body))
		assert.NoError(t, err)
		assert.Equal(t, body, p.event.Body)
		assert.Equal(t, "pull-request", p.channel.Name)
		assert.Equal(t, []string{"test", "webhook", "bridge", "github"}, p.channel.Prefix)
		assert.Len(t, p.event.Attributes, 0)
	})
	t.Run("it should return an error if it cannot unmarshal the event", func(t *testing.T) {
		p := newMockProducer(nil)
		c := NewController(slogger.NewDevNullLogger(), p, "test", "webhook", "bridge")
		body := `not json`
		headers := http.Header{}
		headers.Add(githubEventHeader, "PullRequest")
		err := c.Process(context.Background(), headers, NewMockBody(body))
		assert.ErrorIs(t, err, ErrUnableToEnhanceEvent)
	})
	t.Run("it should publish the body to the correct queue", func(t *testing.T) {
		p := newMockProducer(nil)
		c := NewController(slogger.NewDevNullLogger(), p, "test", "webhook", "bridge")
		body := `{"action":"labelled"}`
		headers := http.Header{}
		headers.Add(githubEventHeader, "PullRequest")
		headers.Add(githubDeliveryHeader, "1234")
		err := c.Process(context.Background(), headers, NewMockBody(body))
		assert.NoError(t, err)
		assert.Equal(t, body, p.event.Body)
		assert.Equal(t, "pull-request", p.channel.Name)
		assert.Equal(t, []string{"test", "webhook", "bridge", "github"}, p.channel.Prefix)
		assert.Contains(t, p.event.Attributes, topic.NewAttribute("action", "labelled"))
		assert.Contains(t, p.event.Attributes, topic.NewAttribute("delivery-id", "1234"))
	})

}
