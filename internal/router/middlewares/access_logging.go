package middlewares

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/cynt4k/router-dns-bridge/internal/router/consts"
	"github.com/cynt4k/router-dns-bridge/internal/router/extension"
	"github.com/cynt4k/router-dns-bridge/pkg/logger"
	"github.com/labstack/echo/v4"
)

func AccessLoggingIgnore() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(consts.KeyIgnoreLogging, true)
			return next(c)
		}
	}
}

func AccessLogging(logger logger.Logger, dev bool) echo.MiddlewareFunc {
	if dev {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				start := time.Now()
				if err := next(c); err != nil {
					c.Error(err)
				}
				stop := time.Now()

				ignoreLogging := c.Get(consts.KeyIgnoreLogging)
				if ignoreLogging != nil {
					if ignoreLogging.(bool) {
						return nil
					}
				}

				req := c.Request()
				res := c.Response()
				logger.Infof("%3d | %s | %s %s %d", res.Status, stop.Sub(start), req.Method, req.URL, res.Size)
				return nil
			}
		}
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			ignoreLogging := c.Get(consts.KeyIgnoreLogging)
			if ignoreLogging != nil {
				if ignoreLogging.(bool) {
					return nil
				}
			}

			req := c.Request()
			res := c.Response()
			type HTTPRequest struct {
				Status        int    `json:"status"`
				RequestMethod string `json:"requestMethod"`
				RequestURL    string `json:"requestUrl"`
				RequestSize   string `json:"requestSize"`
				ResponseSize  int64  `json:"responseSize"`
				UserAgent     string `json:"userAgent"`
				RemoteIP      string `json:"remoteIp"`
				ServerIP      string `json:"serverIp"`
				Referer       string `json:"referer"`
				Latency       string `json:"latency"`
				Protocol      string `json:"protocol"`
			}
			msg := struct {
				RequestID   string      `json:"requestId"`
				HTTPRequest HTTPRequest `json:"httpRequest"`
			}{
				RequestID: extension.GetRequestID(c),
				HTTPRequest: HTTPRequest{
					RequestMethod: req.Method,
					RequestURL:    req.URL.String(),
					RequestSize:   req.Header.Get(echo.HeaderContentLength),
					Status:        res.Status,
					ResponseSize:  res.Size,
					UserAgent:     req.UserAgent(),
					RemoteIP:      c.RealIP(),
					ServerIP:      c.Echo().Server.Addr,
					Referer:       req.Referer(),
					Latency:       strconv.FormatFloat(stop.Sub(start).Seconds(), 'f', 9, 64) + "s",
					Protocol:      req.Proto,
				},
			}
			msgJSON, _ := json.Marshal(&msg)
			logger.Info(string(msgJSON))
			return nil
		}
	}
}
