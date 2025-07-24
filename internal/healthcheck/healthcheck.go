package healthcheck

import (
	"github.com/PULSE-PROXY/pulse-proxy/internal/start"
	"github.com/labstack/echo/v4"
)

// Gateway is the main application instance.
type Gateway struct {
	// private fields for configurations, etc.
	middlewares []echo.MiddlewareFunc
}

// New creates a new Gateway instance.
func New() *Gateway {
	return &Gateway{}
}

func (g *Gateway) Use(middleware ...echo.MiddlewareFunc) {
	g.middlewares = append(g.middlewares, middleware...)
}

// Start starts the gateway server.
func (g *Gateway) Start() {
	start.RunServerLifecycle(g.middlewares)
}
