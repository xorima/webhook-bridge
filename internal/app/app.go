package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/xorima/slogger"
	"github.com/xorima/webhook-bridge/internal/controllers"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Router = chi.Router

type App struct {
	log    *slog.Logger
	port   int
	router Router
}

func NewApp(log *slog.Logger, controller controllers.Controller) *App {
	app := &App{
		port:   3000,
		router: chi.NewRouter(),
		log:    slogger.SubLogger(log, "app"),
	}
	sh := NewSwaggerHandler(app.log)
	hh := NewHealthHandler(app.log)
	app.router.Route("/", func(r Router) {
		sh.RegisterRoutes(r)
		hh.RegisterRoutes(r)
	})
	wh := NewWebhookHandler(app.log, controller)
	app.router.Route("/api", func(r Router) {
		wh.RegisterRoutes(r)
	})
	return app
}

func (a *App) Run() (err error) {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.port),
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      a.router,
	}

	// Start HTTP server.
	srvErr := make(chan error, 1)
	go func() {
		a.log.Info("Starting Server", slog.String("listening", srv.Addr))
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = srv.Shutdown(context.Background())
	return
}
