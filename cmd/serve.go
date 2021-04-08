package cmd

import (
	"fmt"

	"github.com/cynt4k/router-dns-bridge/internal/services"
	"github.com/cynt4k/router-dns-bridge/pkg/logger"
	"github.com/cynt4k/router-dns-bridge/pkg/utils/factory/logger/zap"
	"github.com/labstack/echo/v4"
	"github.com/leandro-lugaresi/hub"
)

type Server struct {
	echo   *echo.Echo
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
		serverURL := fmt.Sprintf("%s:%d", cfg.API.Host, cfg.API.Port)
		server.logger.Info(serverURL)
		if err := server.Start(serverURL); err != nil {
			server.logger.Info("shutting down the server: %w", err)
		}
	}()

	return nil
}

func (s *Server) Start(address string) error {
	return s.echo.Start(address)
}
