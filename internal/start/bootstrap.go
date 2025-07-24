package start

import (
	"fmt"
	"gateway-api/internal/config"
	"gateway-api/internal/logger"
	"gateway-api/internal/middleware"
	"gateway-api/internal/proxy"
	"gateway-api/internal/server"
	yaml_config "gateway-api/internal/yaml"
	"time"

	"github.com/labstack/echo/v4"
)

func StartApp(restartSignal chan<- bool, customMiddlewares ...echo.MiddlewareFunc) server.IServer {
	logger.ClearTerminal()

	server := server.NewEchoAdapter(customMiddlewares...)
	server.HideBanner = true
	server.HidePort = true

	port := yaml_config.GetPort()
	config.PrintBanner(port)

	config.Routes = config.LoadRoutesFromYAML()

	if len(config.Routes) > 0 {
		config.PrintRoutes(config.Routes)
		logger.Info("Routes defined in gateway.yaml: Yes")
	} else {
		logger.Warn("Routes defined in gateway.yaml: No")
	}

	config.ListenerServices(5*time.Second, restartSignal)

	echoInstance := server.Instance().(*echo.Echo)

	echoInstance.Use(middleware.LogRequestMiddleware())
	echoInstance.Use(middleware.CORSMiddleware())

	// Serve index.html if the root route is not defined in gateway.yaml
	if _, ok := config.Routes["/"]; !ok {
		logger.Info("Root route (/) not defined in gateway.yaml. Serving index.html.")
		echoInstance.GET("/", func(c echo.Context) error {
			return c.File("index.html")
		})
	}

	echoInstance.Any("/*", proxy.ReverseProxyHandler)

	go func() {
		if err := server.Start(fmt.Sprintf(":%d", port)); err != nil {
			server.Logger().Error("Server failed to start: %v", err)
		}
	}()

	return server
}
