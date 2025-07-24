package start

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PULSE-PROXY/pulse-proxy/internal/logger"
	"github.com/labstack/echo/v4"
)

// RunServerLifecycle manages the server's lifecycle, including starting, restarting, and graceful shutdown.
func RunServerLifecycle(customMiddlewares []echo.MiddlewareFunc) {
	restartSignal := make(chan bool)
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)

	for {
		logger.Info("Starting PulseProxy server...")
		app := StartApp(restartSignal, customMiddlewares...) // StartApp is already in the 'start' package

		select {
		case <-restartSignal:
			logger.Info("Restart signal received. Shutting down current server...")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			if err := app.Shutdown(ctx); err != nil {
				logger.Error("Server shutdown error: %v", err)
			}
			cancel()
			logger.Info("Server shut down. Restarting...")
			// Loop will continue to start a new server
		case sig := <-quitSignal:
			logger.Info("Quit signal (%s) received. Shutting down PulseProxy...", sig)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			if err := app.Shutdown(ctx); err != nil {
				logger.Error("Server shutdown error: %v", err)
			}
			cancel()
			logger.Info("PulseProxy shut down gracefully.")
			return // Exit the main function
		}
	}
}
