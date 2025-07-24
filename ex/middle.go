package ex

import (
	"pule-proxy/internal/config"
	"pule-proxy/internal/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ApiKeyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	expectedKey := "YOUR_API_KEY_HERE"

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
