package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"pule-proxy/internal/config"
)

func TestPortApp(t *testing.T) {
	// Test case: Default port when no environment variable is set and no YAML config is found
	os.Unsetenv("PORT") // Ensure PORT is not set
	assert.Equal(t, ":9001", config.PortApp(), "Should return default port when no specific port is configured")
}
