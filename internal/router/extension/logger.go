package extension

import (
	"github.com/cynt4k/wygops/pkg/util/random"
	"github.com/labstack/echo/v4"
)

func GetRequestID(c echo.Context) string {
	const randomID = 32
	rid := c.Request().Header.Get(echo.HeaderXRequestID)
	if len(rid) == 0 {
		rid = random.AlphaNumeric(randomID)
	}
	return rid
}
