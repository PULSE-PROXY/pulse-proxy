package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/PULSE-PROXY/pulse-proxy/internal/logger"
	yaml_config "github.com/PULSE-PROXY/pulse-proxy/internal/yaml"
)

var (
	Routes     map[string]string
	RoutesLock sync.RWMutex
)

func ListenerServices(interval time.Duration, restartSignal chan<- bool) {
	go func() {
		var lastModTime time.Time
		var currentPort int

		// Initialize lastModTime and currentPort with the current state of gateway.yaml
		fi, err := os.Stat("gateway.yaml")
		if err == nil {
			lastModTime = fi.ModTime()
			config := yaml_config.LoadConfig()
			if config.Server.Port != nil {
				currentPort = *config.Server.Port
			} else {
				currentPort = 9001 // Default port
			}
		}

		for {
			fi, err := os.Stat("gateway.yaml")
			if err == nil {
				modTime := fi.ModTime()
				if modTime.After(lastModTime) {
					logger.Warn("Change detected in gateway.yaml, reloading configurations...")

					// Reload routes
					RoutesLock.Lock()
					Routes = LoadRoutesFromYAML()
					RoutesLock.Unlock()
					PrintRoutes(Routes)

					// Check for port change
					newConfig := yaml_config.LoadConfig()
					newPort := 9001 // Default port
					if newConfig.Server.Port != nil {
						newPort = *newConfig.Server.Port
					}

					if newPort != currentPort {
						logger.Info("Port changed from %d to %d. Signaling server restart.", currentPort, newPort)
						currentPort = newPort
						restartSignal <- true
					}

					lastModTime = modTime
				}
			}
			time.Sleep(interval)
		}
	}()
}

func PrintRoutes(routes map[string]string) {
	if len(routes) == 0 {
		return
	}

	for path, target := range routes {
		coloredPath := fmt.Sprintf("%s%s%s", logger.ColorYellow, path, logger.ColorReset)
		coloredTarget := fmt.Sprintf("%s%s%s", logger.ColorGreen, target, logger.ColorReset)
		logger.Info("ROUTE: %-7s → %s", coloredPath, coloredTarget)
	}
}
