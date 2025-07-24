package main

import (
	"github.com/PULSE-PROXY/pulse-proxy/ex"
	"github.com/PULSE-PROXY/pulse-proxy/internal/healthcheck"
)

func main() {
	app := healthcheck.New()

	// Middleware de chave de API
	app.Use(ex.ApiKeyMiddleware)

	app.Start()
}