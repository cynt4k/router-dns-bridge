package router

import (
	"github.com/cynt4k/router-dns-bridge/cmd/config"
	"github.com/cynt4k/router-dns-bridge/internal/router/extension"
	"github.com/cynt4k/router-dns-bridge/internal/router/middlewares"
	v1 "github.com/cynt4k/router-dns-bridge/internal/router/v1"
	"github.com/cynt4k/router-dns-bridge/internal/services"
	"github.com/cynt4k/router-dns-bridge/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leandro-lugaresi/hub"
)

type router struct {
	e  *echo.Echo
	v1 *v1.Handlers
}

func New(hub *hub.Hub, config *config.Config, ss *services.Services, logger logger.Logger) *echo.Echo {
	r := newRouter(hub, ss, config, logger.New("router"))

	r.v1.Setup(r.e.Group("/api"))

	return r.e
}

func newEcho(config *config.Config, logger logger.Logger) *echo.Echo {
	const maxAge = 300
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = extension.ErrorHandler(logger)

	e.Use(middlewares.RequestID())
	e.Use(middlewares.AccessLogging(logger.New("access_log"), config.DevMode))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		ExposeHeaders: []string{echo.HeaderXRequestID},
		AllowHeaders:  []string{echo.HeaderContentType, echo.HeaderAuthorization},
		MaxAge:        maxAge,
	}))
	return e
}
