package middleware

import (
	"github.com/PULSE-PROXY/pulse-proxy/internal/logger"
	"github.com/labstack/echo/v4"
)

func LogRequestMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			method := c.Request().Method
			path := c.Request().URL.Path
			ip := c.RealIP()

			logger.Info(
				"%s ACCESS METHOD %s[%s]%s %s- %sPath:%s %s%s %s| %sIP:%s %s%s %s",
				logger.ColorCyan, logger.ColorGreen, method, logger.ColorReset,
				logger.ColorReset, logger.ColorYellow, logger.ColorReset, logger.ColorReset, path, logger.ColorReset,
				logger.ColorYellow, logger.ColorReset, logger.ColorBlue, ip, logger.ColorReset,
			)

			return next(c)
		}
	}
}
