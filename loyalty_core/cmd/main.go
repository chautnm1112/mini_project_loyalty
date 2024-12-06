package main

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"loyalty_core/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var cfg *config.Config
var logger *zap.Logger

func main() {
	cfg := config.LoadConfig()

	err := serverAction(cfg)
	if err != nil {
		return
	}
}

func serverAction(cfg *config.Config) error {
	var err error
	logger, err = zap.NewProduction()

	if err != nil {
		logger.Error("Error init zap logger", zap.Error(err))
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			w.Write([]byte("Hello, World!"))
		}
	})
	//_, err = newService(cfg)
	//if err != nil {
	//	logger.Error("Error init servers", zap.Error(err))
	//}

	startServer(mux, cfg)
	return nil
}

func startServer(handler http.Handler, cfg *config.Config) {
	if cfg == nil {
		logger.Fatal("Config is nil")
	}

	if cfg.Host == "" {
		logger.Fatal("Host in config is empty")
	}

	server := &http.Server{
		Addr:    ":" + cfg.Host,
		Handler: handler,
	}

	timeWait := 15 * time.Second

	signChan := make(chan os.Signal, 1)

	go func() {
		logger.Info("Starting server on :" + cfg.Host)
		if err := server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			logger.Error("Could not listen on "+server.Addr+": ", zap.Error(err))
		}
	}()

	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
	<-signChan

	logger.Info("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), timeWait)
	defer func() {
		logger.Info("Close another connection")
		cancel()
	}()

	logger.Info("Stop http server")

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown: ", zap.Error(err))
	}

	close(signChan)
	logger.Info("Server stopped gracefully")
}
