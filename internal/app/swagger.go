package app

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/xorima/webhook-bridge/docs"
	"log/slog"
	"net/http"
)

const swaggerPath = "/swagger"

type SwaggerHandler struct {
	log *slog.Logger
}

func NewSwaggerHandler(log *slog.Logger) *SwaggerHandler {
	return &SwaggerHandler{
		log: log.With(slog.String("handler", "swagger")),
	}
}

func (sh *SwaggerHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	sh.log.InfoContext(r.Context(), fmt.Sprintf("redirecting to %s/", swaggerPath), slog.String("path", r.URL.Path))
	w.Header().Set("Location", fmt.Sprintf("%s/", swaggerPath))
	w.WriteHeader(http.StatusFound)
}

func (sh *SwaggerHandler) RegisterRoutes(r Router) {
	r.Mount(swaggerPath, httpSwagger.WrapHandler)
	r.Get("/", sh.Redirect)
	r.Get("/swagger-ui", sh.Redirect)
	r.Get(swaggerPath, sh.Redirect)
}
