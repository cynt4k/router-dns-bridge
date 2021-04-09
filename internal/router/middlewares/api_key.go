package middlewares

import (
	"fmt"
	"net/http"

	"github.com/cynt4k/router-dns-bridge/internal/router/consts"
	"github.com/labstack/echo/v4"
)

func AcessControlAPIKey(key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			providedKey := req.Header.Get(consts.HeaderAPIKey)

			if providedKey != key {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					fmt.Sprintf("you are not permitted to request to '%s'", req.URL.Path),
				)
			}
			return next(c)
		}
	}
}
