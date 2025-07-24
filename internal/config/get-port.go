package config

import (
	"strconv"

	yaml_config "github.com/PULSE-PROXY/pulse-proxy/internal/yaml"
)

func PortApp() string {
	config := yaml_config.LoadConfig()

	var port int

	if config.Server.Port != nil {
		port = *config.Server.Port
	} else {
		port = 9001
	}

	str := ":" + strconv.Itoa(port)

	return str
}
