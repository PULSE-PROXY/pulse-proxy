package main

import (
	"pule-proxy/ex"
	"pule-proxy/internal/healthcheck"

)

func main() {
	app := healthcheck.New()

	// Middleware de chave de API
	app.Use(ex.ApiKeyMiddleware)

	app.Start()
}