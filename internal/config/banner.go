package config

import (
	"fmt"
	"gateway-api/internal/logger"
)

func PrintBanner(port int) {

	portText := fmt.Sprintf("http://localhost:%d", port)

	clickablePort := fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", portText, logger.ColorYellow+portText+logger.ColorReset)

	fmt.Println(logger.ColorGreen + `
 _____       _            _____                     
|  __ \     | |          |  __ \                    
| |__) |   _| |___  ___  | |__) | __ _____  ___   _ 
|  ___/ | | | / __|/ _ \ |  ___/ '__/ _ \ \/ / | | |
| |   | |_| | \__ \  __/ | |   | | | (_) >  <| |_| |
|_|    \__,_|_|___/\___| |_|   |_|  \___/_/\_\\__, |
                                               __/ |
                                               |___/ 

⇨ PulseProxy started on port: ` + clickablePort + logger.ColorReset + `
`)
}
