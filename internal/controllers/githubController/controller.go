package githubController

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xorima/slogger"
	"github.com/xorima/webhook-bridge/internal/data/topic"
	"github.com/xorima/webhook-bridge/internal/infrastructure/errs"
	"io"
	"log/slog"
	"net/http"
	"unicode"
)

const (
	githubEventHeader    = "X-GitHub-Event"
	githubDeliveryHeader = "X-GitHub-Delivery"
)

var (
	ErrMissingHeader        = fmt.Errorf("missing %s header from request", githubEventHeader)
	ErrCannotReadBody       = errors.New("unable to read body")
	ErrFailedToPublish      = errors.New("unable to publish message onto producer")
	ErrUnableToEnhanceEvent = errors.New("unable to enhance event with additional attributes")
)

type Controller struct {
	log      *slog.Logger
	prefix   []string // prefix is any prefix we want on channels, useful for multi env
	producer topic.EventProducer
}

func NewController(log *slog.Logger, producer topic.EventProducer, prefix ...string) *Controller {
	return &Controller{
		log:      slogger.SubLogger(log, "githubController"),
		prefix:   prefix,
		producer: producer,
	}
}

func (c *Controller) Process(ctx context.Context, header http.Header, body io.ReadCloser) error {
	e := header.Get("X-GitHub-Event")
	if e == "" {
		c.log.WarnContext(ctx, "header empty", slog.String("header", "X-GitHub-Event"))
		return ErrMissingHeader
	}
	b, err := io.ReadAll(body)
	if err != nil {
		c.log.WarnContext(ctx, "unable to parse body", slogger.ErrorAttr(err))
		return errs.WrapError(err, ErrCannotReadBody)
	}
	return c.githubEventTopics(ctx, e, string(b), header)
}

func (c *Controller) githubEventTopics(ctx context.Context, event, body string, header http.Header) error {
	// the purpose of this is to fan out events to correct queues based on the topic names.
	name := pascalToHyphen(event)
	chn := topic.NewChannel(name).WithPrefix(c.prefix...).WithPrefix("github")
	attr, err := c.enhanceEvent(ctx, event, body, header)
	if err != nil {
		return err
	}
	evt := topic.NewEvent("1.0.0", body, attr...)
	err = c.producer.Produce(ctx, chn, evt)
	if err != nil {
		c.log.ErrorContext(ctx, ErrFailedToPublish.Error(), slogger.ErrorAttr(err))
		return errs.WrapError(err, ErrFailedToPublish)
	}
	return nil
}

func pascalToHyphen(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			result = append(result, '-')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

// enhanceEvent is responsible for adding in attributes/metadata for the event which could be useful
func (c *Controller) enhanceEvent(ctx context.Context, event, body string, headers http.Header) ([]topic.Attribute, error) {
	var resp []topic.Attribute
	var data map[string]any
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		return nil, errs.WrapError(err, ErrUnableToEnhanceEvent)
	}
	action, exists := data["action"]
	if exists {
		resp = append(resp, topic.NewAttribute("action", action))
	}
	id := headers.Get("X-GitHub-Delivery")
	if id != "" {
		resp = append(resp, topic.NewAttribute("delivery-id", id))
	}
	return resp, nil
}
