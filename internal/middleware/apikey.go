package middleware

import (
	"net/http"

	"github.com/PULSE-PROXY/pulse-proxy/internal/config"
	"github.com/PULSE-PROXY/pulse-proxy/internal/logger"
	"github.com/labstack/echo/v4"
)

func APIKeyMiddleware(expectedKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == http.MethodOptions {
				return next(c)
			}
			key := c.Request().Header.Get("x-api-key")
			if key != expectedKey {
				logger.Warn("Invalid API Key: %s", key)
				return config.JsonResponse(c, http.StatusUnauthorized, "Invalid API key", nil, nil)
			}
			return next(c)
		}
	}
}
