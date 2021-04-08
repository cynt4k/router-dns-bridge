package router

import (
	"github.com/cynt4k/router-dns-bridge/cmd/config"
	v1 "github.com/cynt4k/router-dns-bridge/internal/router/v1"
	"github.com/cynt4k/router-dns-bridge/internal/services"
	"github.com/cynt4k/router-dns-bridge/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/leandro-lugaresi/hub"
)

type router struct {
	e  *echo.Echo
	v1 *v1.Handlers
}

func New(hub *hub.Hub, config *config.Config, ss *services.Services, logger logger.Logger) *echo.Echo {
	r := newRouter(hub, ss, config, logger)

	r.v1.Setup(r.e.Group("/api"))

	return r.e
}

func newEcho(config *config.Config, logger logger.Logger) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true
	return e
}
