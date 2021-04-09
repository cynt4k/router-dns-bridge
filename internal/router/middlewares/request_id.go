package middlewares

import (
	"github.com/cynt4k/router-dns-bridge/internal/router/extension"
	"github.com/labstack/echo/v4"
)

func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderXRequestID, extension.GetRequestID(c))
			return next(c)
		}
	}
}
