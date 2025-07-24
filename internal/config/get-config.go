package config

import (
	"pule-proxy/internal/logger"
	yaml_config "pule-proxy/internal/yaml"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadRoutesFromYAML() map[string]string {
	filePath := "gateway.yaml"

	content, err := os.ReadFile(filePath)
	if err != nil {
		logger.Warn("File %s not found. No routes will be loaded.", filePath)
		return map[string]string{}
	}

	var config yaml_config.GatewayConfig
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		logger.Error("Error parsing YAML: %v", err)
		return map[string]string{}
	}

	if len(config.Routes) == 0 {
		logger.Warn("No route defined in file %s.", filePath)
	}

	routesMap := make(map[string]string)
	for _, route := range config.Routes {
		if route.Path != "" && route.URI != "" {
			routesMap[route.Path] = route.URI
		}
	}

	return routesMap
}
