package middleware

import (
	"gateway-api/internal/logger"
	yaml_config "gateway-api/internal/yaml"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CORSMiddleware() echo.MiddlewareFunc {
	corsCfg := yaml_config.GlobalConfig.GlobalCORS

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			origin := c.Request().Header.Get("Origin")

			// ORIGINS
			if containsWildcard(corsCfg.AllowedOrigins) {
				c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			} else {
				for _, allowed := range corsCfg.AllowedOrigins {
					if strings.EqualFold(origin, allowed) {
						c.Response().Header().Set("Access-Control-Allow-Origin", origin)
						if corsCfg.AllowCredentials {
							c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
						}
						c.Response().Header().Set("Vary", "Origin")
						break
					}
				}
			}

			// HEADERS
			if containsWildcard(corsCfg.AllowedHeaders) {
				c.Response().Header().Set("Access-Control-Allow-Headers", "*")
			} else {
				c.Response().Header().Set("Access-Control-Allow-Headers", strings.Join(corsCfg.AllowedHeaders, ", "))
			}

			// METHODS
			if containsWildcard(corsCfg.AllowedMethods) {
				c.Response().Header().Set("Access-Control-Allow-Methods", "*")
			} else {
				c.Response().Header().Set("Access-Control-Allow-Methods", strings.Join(corsCfg.AllowedMethods, ", "))
			}

			// OPTIONS (preflight)
			if c.Request().Method == http.MethodOptions {
				return c.NoContent(http.StatusNoContent)
			}

			if containsWildcard(corsCfg.AllowedOrigins) && corsCfg.AllowCredentials {
				logger.Warn("CORS: Using '*' in 'allowedOrigins' with 'allowCredentials: true' may violate the CORS specification.")
			}

			return next(c)
		}
	}
}

func containsWildcard(list []string) bool {
	for _, v := range list {
		if v == "*" {
			return true
		}
	}
	return false
}
