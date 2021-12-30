package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/daveseco7/http-echo-server/internal/config"
	"github.com/daveseco7/http-echo-server/internal/service"
	transportHTTP "github.com/daveseco7/http-echo-server/internal/transport/http"
)

func main() {
	// initialize root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, _ := zap.NewProduction()
	// flushes buffer, if any
	defer logger.Sync() // nolint:errcheck
	sugar := logger.Sugar()

	// load configuration
	conf, err := config.GetConfig()
	if err != nil {
		sugar.Fatalf("fail loading configuration: %e", err)
	}

	// setup service
	servConfig := conf.WorkloadConfig()
	serv := service.NewService(servConfig)

	// setup handler
	httpServerConfig := conf.HTTPServerConfig()
	server := &http.Server{
		ReadTimeout:  httpServerConfig.ReadTimeout * time.Second,
		WriteTimeout: httpServerConfig.WriteTimeout * time.Second,
		IdleTimeout:  httpServerConfig.IdleTimeout * time.Second,
		Addr:         ":8080",
		Handler:      transportHTTP.New(serv, sugar),
	}

	// handle http servers lifecycle
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// start http server
	go func() {
		err := server.ListenAndServe()
		if err == http.ErrServerClosed {
			sugar.Infof("http server terminated: %e", err)
		} else {
			sugar.Infof("http server terminated unexpectedly: %e", err)
		}
		sigs <- syscall.SIGQUIT
	}()

	sig := <-sigs
	switch sig {
	// Gracefully shutdown the service
	case os.Interrupt, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
		// Shutdown http server
		err := server.Shutdown(ctx)
		if err != nil {
			sugar.Infof("error shutting down http server: %e", err)
		}
	}
}
