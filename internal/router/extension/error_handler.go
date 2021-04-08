package extension

import (
	"fmt"
	"net/http"

	"github.com/cynt4k/router-dns-bridge/pkg/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ErrorMessage struct {
	Code    int
	Message string
}

func ErrorHandler(logger logger.Logger) echo.HTTPErrorHandler {
	return func(e error, c echo.Context) {
		var (
			code int
			body interface{}
		)

		switch err := e.(type) {
		case nil:
			return
		case *echo.HTTPError:
			if err.Internal != nil {
				if herr, ok := err.Internal.(*echo.HTTPError); ok {
					err = herr
				}
			}

			code = err.Code

			switch m := err.Message.(type) {
			case string:
				body = echo.Map{
					"code":    code,
					"message": m,
				}
			case error:
				body = echo.Map{
					"code":    code,
					"message": m.Error(),
				}
			default:
				body = echo.Map{
					"code":    code,
					"message": m,
				}
			}
		default:
			logger.Error(err.Error())
			code = http.StatusInternalServerError
			body = echo.Map{
				"code":    code,
				"message": fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err.Error()),
			}
		}

		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				e = c.NoContent(code)
			} else {
				e = json(c, code, body, jsoniter.ConfigFastest)
			}
			if e != nil {
				logger.Warn("failed so send error response", zap.Error(e))
			}
		}
	}
}
