package yaml_config

func DefaultCORS() GatewayConfig {
	return GatewayConfig{
		GlobalCORS: GlobalCORS{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"OPTIONS",
			},
			AllowedHeaders: []string{
				"*",
			},
			AllowCredentials: false,
		},
	}
}
