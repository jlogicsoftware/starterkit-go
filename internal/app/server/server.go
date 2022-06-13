package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jlogicsoftware/starterkit/internal/pkg/config"
	"github.com/jlogicsoftware/starterkit/internal/pkg/logger"
)

type Server struct {
	*http.Server
}

func New() (*Server, error) {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// r.Get("/swagger/*", http.StripPrefix("/swagger/", http.FileServer(http.Dir("/swagger"))))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", config.Get().Port),
		Handler:           r,
		TLSConfig:         &tls.Config{},
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    0,
		TLSNextProto:      map[string]func(*http.Server, *tls.Conn, http.Handler){},
		ConnState: func(net.Conn, http.ConnState) {
		},
		ErrorLog: log.New(&logger.Logger{}, "", 0),
		BaseContext: func(net.Listener) context.Context {
			return context.Background()
		},
		ConnContext: func(ctx context.Context, c net.Conn) context.Context {
			return ctx
		},
	}
	return &Server{&srv}, nil
}

func Start(srv *Server) error {
	logger.Info().Msg("Starting server at " + srv.Addr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}
