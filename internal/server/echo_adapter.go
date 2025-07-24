package server

import (
	"context"

	"github.com/PULSE-PROXY/pulse-proxy/internal/logger"
	"github.com/labstack/echo/v4"
)

// EchoAdapter is an adapter for the Echo framework that implements the IServer interface.
type EchoAdapter struct {
	*echo.Echo
	logger      logger.ILogger
	middlewares []echo.MiddlewareFunc
}

func (e *EchoAdapter) Use(middleware ...echo.MiddlewareFunc) {
	e.middlewares = append(e.middlewares, middleware...)
}

// NewEchoAdapter creates a new instance of the EchoAdapter.
func NewEchoAdapter(middlewares ...echo.MiddlewareFunc) *EchoAdapter {
	adapter := &EchoAdapter{echo.New(), logger.NewLogger(), middlewares}
	for _, m := range adapter.middlewares {
		adapter.Echo.Use(m)
	}
	return adapter
}

func (e *EchoAdapter) Start(address string) error {
	return e.Echo.Start(address)
}

func (e *EchoAdapter) Shutdown(ctx context.Context) error {
	return e.Echo.Shutdown(ctx)
}

// Logger returns the logger instance associated with the server.
func (e *EchoAdapter) Logger() logger.ILogger {
	return e.logger
}

// Instance returns the underlying server instance (*echo.Echo).
func (e *EchoAdapter) Instance() interface{} {
	return e.Echo
}
