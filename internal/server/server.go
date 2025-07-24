package server

import (
	"context"

	"github.com/PULSE-PROXY/pulse-proxy/internal/logger"
)

// IServer is the interface that abstracts the web server.
// It allows us to swap out the underlying framework (e.g., Echo) without
// changing the application's entry point.
type IServer interface {
	Start(address string) error
	Shutdown(ctx context.Context) error
	// Logger returns the logger instance associated with the server.
	// We need a generic logger interface here eventually, but for now,
	// we'll use Echo's to minimize changes.
	Logger() logger.ILogger
	// Instance returns the underlying server instance (e.g., *echo.Echo)
	// This is a temporary escape hatch to allow the start package to
	// configure routes and middleware on the specific framework instance.
	Instance() interface{}
}
