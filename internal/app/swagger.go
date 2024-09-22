package app

import (
	_ "github-bridge/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type SwaggerHandler struct {
}

func NewSwaggerHandler() *SwaggerHandler {
	return &SwaggerHandler{}
}

func (sh *SwaggerHandler) RegisterRoutes(r Router) {
	r.Mount("/swagger", httpSwagger.WrapHandler)

}
