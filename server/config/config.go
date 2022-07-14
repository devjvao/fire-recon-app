package config

import (
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	Port     = "PORT"
	LogLevel = "LOG_LEVEL"
)

var instance *Config
var once sync.Once

// GetConfig creates a single instance of the configuration object.
func GetConfig() *Config {
	once.Do(func() {
		env, err := godotenv.Read()
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		path := filepath.Join(dir, "logs")
		instance = &Config{
			Port:     "7070",
			LogPath:  path,
			LogLevel: log.ErrorLevel,
		}
		if err == nil {
			if env[Port] != "" {
				instance.Port = env[Port]
			}
			if env[LogLevel] != "" {
				instance.LogLevel, _ = log.ParseLevel(env[LogLevel])
			}
		}
		if _, err := os.Stat(instance.LogPath); os.IsNotExist(err) {
			_ = os.Mkdir(instance.LogPath, 0755)
		}
	})
	return instance
}

// parseUint converts a string to uint and returns the defaultValue in case of conversion error.
func parseUint(value string, defaultValue uint) uint {
	if parsed, err := strconv.ParseUint(value, 10, 64); err == nil {
		return uint(parsed)
	}
	return defaultValue
}
