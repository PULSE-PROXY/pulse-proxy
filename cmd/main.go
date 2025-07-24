package main

import (
	"gateway-api/ex"
	"gateway-api/internal/healthcheck"

)

func main() {
	app := healthcheck.New()

	// Middleware de chave de API
	app.Use(ex.ApiKeyMiddleware)

	app.Start()
}