package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/lestrrat-go/server-starter/listener"
	"golang.org/x/sync/errgroup"
)

//go:generate wire ./...

func NewListener() (net.Listener, error) {
	return net.Listen("tcp", "127.0.0.1:8080")
}

func NewServerStarterListener() (net.Listener, error) {
	listeners, err := listener.ListenAll()
	if err != nil {
		return nil, err
	}
	if len(listeners) == 0 {
		return nil, err
	}
	return listeners[0], nil
}

func NewServer() (*http.Server, error) {
	r := chi.NewRouter()
	r.Mount("/", NewAssetsHandler())
	r.Route("/api", func(r chi.Router) {
		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome"))
		})
	})
	return &http.Server{Handler: r}, nil
}

type App struct {
	listener net.Listener
	server   *http.Server
}

func (a *App) Start() error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		if err := a.server.Serve(a.listener); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	eg.Go(func() error {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGTERM, os.Interrupt)
		<-signalChan

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		return a.server.Shutdown(ctx)
	})
	return eg.Wait()
}

func start() error {
	app, cleanup, err := InitializeApp()
	if err != nil {
		return err
	}
	defer cleanup()

	return app.Start()
}

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}
