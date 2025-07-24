package proxy

import (
	"fmt"
	"gateway-api/internal/config"
	"gateway-api/internal/logger"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

func ReverseProxyHandler(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("PANIC GATEWAY: %-7s", r)
			_ = config.JsonResponse(c, http.StatusInternalServerError, "Internal proxy error", nil, nil)
		}
	}()

	config.RoutesLock.RLock()
	defer config.RoutesLock.RUnlock()

	path := c.Request().URL.Path

	for prefix, target := range config.Routes {
		if strings.HasPrefix(path, prefix) && target != "" {
			targetURL, _ := url.Parse(target)
			finalPath := path
			originalRequest := c.Request()

			proxy := httputil.NewSingleHostReverseProxy(targetURL)

			proxy.Director = func(req *http.Request) {
				req.URL.Scheme = targetURL.Scheme
				req.URL.Host = targetURL.Host
				req.URL.Path = finalPath
				req.URL.RawQuery = originalRequest.URL.RawQuery
				req.Host = targetURL.Host
				req.Header = originalRequest.Header.Clone()
			}

			proxy.ModifyResponse = func(resp *http.Response) error {
				for _, cookie := range resp.Header.Values("Set-Cookie") {
					logger.Info("Set-Cookie: %-7s", cookie)
				}
				return nil
			}

			proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
				coloredPath := fmt.Sprintf("%s%s%s", logger.ColorYellow, err.Error(), logger.ColorReset)
				logger.Error("SERVICE UNAVAILABLE: %-7s", coloredPath)
				if !c.Response().Committed {
					_ = config.JsonResponse(c, http.StatusBadGateway, "Service Unavailable", nil, nil)
				}
			}

			proxy.ServeHTTP(c.Response(), originalRequest)
			return nil
		}
	}

	coloredPath := fmt.Sprintf("%s%s%s", logger.ColorRed, path, logger.ColorRed)
	logger.Error("SERVICE OR ROUE NOT FOUND: %-7s", coloredPath)
	return config.JsonResponse(c, http.StatusNotFound, "Service or route not found", nil, nil)
}
