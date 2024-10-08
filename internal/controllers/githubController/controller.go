package githubController

import (
	"context"
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

const githubEventHeader = "X-GitHub-Event"

var (
	ErrMissingHeader   = fmt.Errorf("missing %s header from request", githubEventHeader)
	ErrCannotReadBody  = errors.New("unable to read body")
	ErrFailedToPublish = errors.New("unable to publish message onto producer")
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
	return c.githubEventTopics(ctx, e, string(b))
}

func (c *Controller) githubEventTopics(ctx context.Context, event, body string) error {
	// the purpose of this is to fan out events to correct queues based on the topic names.
	name := pascalToHyphen(event)
	chn := topic.NewChannel(name).WithPrefix(c.prefix...).WithPrefix("github")
	evt := topic.NewEvent("1.0.0", body)
	err := c.producer.Produce(ctx, chn, evt)
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
