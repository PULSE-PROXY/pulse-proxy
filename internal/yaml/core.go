package yaml_config

type GatewayConfig struct {
	Server struct {
		Port *int `yaml:"port"`
	} `yaml:"server"`

	Routes []Route `yaml:"routes"`

	GlobalCORS GlobalCORS `yaml:"globalCors"`
}

type GlobalCORS struct {
	AllowedOrigins   []string `yaml:"allowed_origins"`
	AllowedMethods   []string `yaml:"allowed_methods"`
	AllowedHeaders   []string `yaml:"allowed_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
}

type Route struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	URI  string `yaml:"uri"`
}
