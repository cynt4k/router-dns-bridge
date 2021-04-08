package v1

import (
	"net/http"

	"github.com/cynt4k/router-dns-bridge/internal/services/powerdns"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	PowerDNS powerdns.PowerDNS
}

func (h *Handlers) Setup(e *echo.Group) {
	apiNoAuth := e.Group("/v1")
	{
		apiNoAuth.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, http.StatusText(http.StatusOK)) })
	}
}
