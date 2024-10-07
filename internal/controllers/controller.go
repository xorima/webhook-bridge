package controllers

import (
	"context"
	"io"
	"net/http"
)

type Controller interface {
	Process(ctx context.Context, header http.Header, body io.ReadCloser) error
}
