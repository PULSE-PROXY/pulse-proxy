package config

import (
	yaml_config "gateway-api/internal/yaml"
	"strconv"
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
