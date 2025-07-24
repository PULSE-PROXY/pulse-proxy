package config_test

import (
	"os"
	"testing"

	"github.com/PULSE-PROXY/pulse-proxy/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadRoutesFromYAML(t *testing.T) {
	// Create a temporary gateway.yaml file for testing
	tempYAMLContent := `
server:
  port: 8080
routes:
  - path: /api/users
    uri: http://localhost:3000/users
  - path: /api/products
    uri: http://localhost:4000/products
`

	err := os.WriteFile("gateway.yaml", []byte(tempYAMLContent), 0644)
	assert.NoError(t, err, "Should write temporary gateway.yaml without error")
	defer os.Remove("gateway.yaml") // Clean up the temporary file

	// Test case: Routes loaded correctly from YAML
	expectedRoutes := map[string]string{
		"/api/users":    "http://localhost:3000/users",
		"/api/products": "http://localhost:4000/products",
	}
	actualRoutes := config.LoadRoutesFromYAML()
	assert.Equal(t, expectedRoutes, actualRoutes, "Should load routes correctly from gateway.yaml")

	// Test case: No gateway.yaml file (should return empty map)
	os.Remove("gateway.yaml") // Remove the file to simulate not found
	actualRoutes = config.LoadRoutesFromYAML()
	assert.Empty(t, actualRoutes, "Should return empty map when gateway.yaml is not found")
}
