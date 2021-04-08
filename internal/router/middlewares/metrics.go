package middlewares

import (
	"strconv"

	"github.com/cynt4k/router-dns-bridge/internal/router/consts"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var requestCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace: "wygops",
	Name:      "http_requests_total",
}, []string{"code", "method"})

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "wygops",
	Name:      "http_duration_seconds",
}, []string{"path"})

func MetricsIgnore() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(consts.KeyIgnoreMetrics, true)
			return next(c)
		}
	}
}

func RequestCounter() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			e := next(c)

			if ignoreMetrics := c.Get(consts.KeyIgnoreMetrics); ignoreMetrics != nil {
				if ignoreMetrics.(bool) {
					return nil
				}
			}
			var code int
			if e != nil {
				switch err := e.(type) {
				case *echo.HTTPError:
					code = err.Code
				default:
					code = 500
				}
			} else {
				code = c.Response().Status
			}
			requestCounter.WithLabelValues(strconv.Itoa(code), c.Request().Method).Inc()
			return e
		}
	}
}

func HTTPDuration() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL
			timer := prometheus.NewTimer(httpDuration.WithLabelValues(path.Path))
			e := next(c)

			if ignoreMetrics := c.Get(consts.KeyIgnoreMetrics); ignoreMetrics != nil {
				if ignoreMetrics.(bool) {
					return nil
				}
			}

			timer.ObserveDuration()
			return e
		}
	}
}
