package config

import (
	"time"

	"github.com/PULSE-PROXY/pulse-proxy/internal/server"
	"github.com/labstack/echo/v4"
)

func JsonResponse(c echo.Context, status int, detail string, content interface{}, err interface{}) error {
	return c.JSON(status, server.Response{
		Status:  status,
		Detail:  detail,
		Content: content,
		Error:   err,
		Metadata: server.Metadata{
			Service:   "gateway",
			Version:   "1.0.0",
			Timestamp: time.Now(),
		},
	})
}
