package app

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Router = chi.Router

type App struct {
	port   int
	router Router
}

func NewApp() *App {
	app := &App{
		port:   3000,
		router: chi.NewRouter(),
	}
	sh := NewSwaggerHandler()
	app.router.Route("/", func(r Router) {
		sh.RegisterRoutes(r)
	})
	wh := NewWebhookHandler()
	app.router.Route("/api", func(r Router) {
		wh.RegisterRoutes(r)
	})
	return app
}
func (a *App) Run() (err error) {
	return http.ListenAndServe(":3000", a.router)
}
