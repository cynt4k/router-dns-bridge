package cmd

import (
	"fmt"
	"os"

	"github.com/cynt4k/router-dns-bridge/internal/services"
	"github.com/cynt4k/router-dns-bridge/pkg/logger"
	"github.com/cynt4k/router-dns-bridge/pkg/utils/factory/logger/zap"
	"github.com/leandro-lugaresi/hub"
)

type Server struct {
	logger logger.Logger
	SS     *services.Services
	Hub    *hub.Hub
}

func Serve() error {
	hub := hub.New()
	zapLogger := zap.New(cfg)
	logger.SetLogger(zapLogger)

	zapLogger.Info("creating server")
	server, err := newServer(hub, zapLogger, cfg)

	if err != nil {
		return fmt.Errorf("error while creating server: %w", err)
	}

	go func() {
		ch := make(chan bool)
		for {
			server.logger.Info("serving server")
			<-ch
			os.Exit(0)
		}
	}()

	return nil
}
