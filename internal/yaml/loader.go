package yaml_config

import (
	"pule-proxy/internal/logger"
	"os"

	"gopkg.in/yaml.v2"
)

var GlobalConfig GatewayConfig

func getDefaultConfig() GatewayConfig {
	return GatewayConfig{
		GlobalCORS: DefaultCORS().GlobalCORS,
	}
}



func LoadGatewayConfig() {
	GlobalConfig = getDefaultConfig() // Initialize with defaults

	filePath := "gateway.yaml"

	content, err := os.ReadFile(filePath)
	if err != nil {
		logger.Warn("⚠️ Arquivo %s não encontrado. Aplicando configurações padrão.", filePath)
		return // Defaults are already set
	}

	err = yaml.Unmarshal(content, &GlobalConfig)
	if err != nil {
		logger.Error("❌ Erro ao fazer parse do YAML: %v", err)
		GlobalConfig = getDefaultConfig() // Revert to defaults on parse error
		return
	}

	if GlobalConfig.Server.Port != nil {
		logger.Info("Loaded port from gateway.yaml: %d", *GlobalConfig.Server.Port)
	} else {
		logger.Info("Loaded port from gateway.yaml: (default) 9001")
	}
}

func LoadConfig() *GatewayConfig {
	LoadGatewayConfig()
	return &GlobalConfig
}

func GetPort() int {
	config := LoadConfig()
	if config.Server.Port != nil {
		return *config.Server.Port
	}
	return 9001
}

func GetRoutes() []Route {
	config := LoadConfig()
	return config.Routes
}

func GetGlobalCORS() GlobalCORS {
	config := LoadConfig()
	return config.GlobalCORS
}
