package yaml_config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gateway-api/internal/yaml"
)

func TestLoadGatewayConfig_ValidYAML(t *testing.T) {
	yaml_config.GlobalConfig = yaml_config.GatewayConfig{} // Reset global config
	// Create a temporary valid gateway.yaml file
	tempYAMLContent := `
server:
  port: 8080
routes:
  - path: /test
    uri: http://test.com
global_cors:
  allowed_origins:
    - http://example.com
  allowed_methods:
    - GET
  allowed_headers:
    - Content-Type
`

	err := os.WriteFile("gateway.yaml", []byte(tempYAMLContent), 0644)
	assert.NoError(t, err, "Should write temporary gateway.yaml without error")
	defer os.Remove("gateway.yaml") // Clean up

	yaml_config.LoadGatewayConfig()

	assert.Equal(t, 8080, *yaml_config.GlobalConfig.Server.Port, "Should load correct port")
	assert.Len(t, yaml_config.GlobalConfig.Routes, 1, "Should load correct number of routes")
	assert.Equal(t, "/test", yaml_config.GlobalConfig.Routes[0].Path, "Should load correct route path")
	assert.Equal(t, "http://test.com", yaml_config.GlobalConfig.Routes[0].URI, "Should load correct route URI")
	assert.Contains(t, yaml_config.GlobalConfig.GlobalCORS.AllowedOrigins, "http://example.com", "Should load correct CORS origin")
}

func TestLoadGatewayConfig_MissingYAML(t *testing.T) {
	yaml_config.GlobalConfig = yaml_config.GatewayConfig{} // Reset global config
	// Ensure gateway.yaml does not exist
	os.Remove("gateway.yaml")

	yaml_config.LoadGatewayConfig()

	// Should apply default config, including default CORS
	assert.NotNil(t, yaml_config.GlobalConfig.GlobalCORS, "GlobalCORS should not be nil")
	assert.Contains(t, yaml_config.GlobalConfig.GlobalCORS.AllowedOrigins, "*", "Should apply default CORS origins")
}

func TestLoadGatewayConfig_InvalidYAML(t *testing.T) {
	yaml_config.GlobalConfig = yaml_config.GatewayConfig{} // Reset global config
	// Create a temporary invalid gateway.yaml file
	tempYAMLContent := `
invalid: yaml: ---
`

	err := os.WriteFile("gateway.yaml", []byte(tempYAMLContent), 0644)
	assert.NoError(t, err, "Should write temporary invalid gateway.yaml without error")
	defer os.Remove("gateway.yaml") // Clean up

	yaml_config.LoadGatewayConfig()

	// Should apply default config due to parsing error
	assert.NotNil(t, yaml_config.GlobalConfig.GlobalCORS, "GlobalCORS should not be nil")
	assert.Contains(t, yaml_config.GlobalConfig.GlobalCORS.AllowedOrigins, "*", "Should apply default CORS origins on invalid YAML")
}

func TestLoadGatewayConfig_DefaultCORS(t *testing.T) {
	yaml_config.GlobalConfig = yaml_config.GatewayConfig{} // Reset global config
	// Create a temporary gateway.yaml with no CORS specified
	tempYAMLContent := `
server:
  port: 8080
`

	err := os.WriteFile("gateway.yaml", []byte(tempYAMLContent), 0644)
	assert.NoError(t, err, "Should write temporary gateway.yaml without error")
	defer os.Remove("gateway.yaml") // Clean up

	yaml_config.LoadGatewayConfig()

	assert.NotNil(t, yaml_config.GlobalConfig.GlobalCORS, "GlobalCORS should not be nil")
	assert.Contains(t, yaml_config.GlobalConfig.GlobalCORS.AllowedOrigins, "*", "Should apply default CORS origins when not specified")
	assert.Contains(t, yaml_config.GlobalConfig.GlobalCORS.AllowedMethods, "GET", "Should apply default CORS methods when not specified")
}

func TestGetters(t *testing.T) {
	yaml_config.GlobalConfig = yaml_config.GatewayConfig{} // Reset global config
	// Create a temporary valid gateway.yaml file for getters
	tempYAMLContent := `
server:
  port: 9000
routes:
  - path: /api
    uri: http://api.com
global_cors:
  allowed_origins:
    - http://test.com
`

	err := os.WriteFile("gateway.yaml", []byte(tempYAMLContent), 0644)
	assert.NoError(t, err, "Should write temporary gateway.yaml without error")
	defer os.Remove("gateway.yaml") // Clean up

	yaml_config.LoadGatewayConfig()

	assert.Equal(t, 9000, yaml_config.GetPort(), "GetPort should return correct port")
	assert.Len(t, yaml_config.GetRoutes(), 1, "GetRoutes should return correct number of routes")
	assert.Equal(t, "/api", yaml_config.GetRoutes()[0].Path, "GetRoutes should return correct route path")
	assert.Contains(t, yaml_config.GetGlobalCORS().AllowedOrigins, "http://test.com", "GetGlobalCORS should return correct CORS origins")
}
