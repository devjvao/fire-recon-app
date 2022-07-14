package config

import log "github.com/sirupsen/logrus"

// Config is the configuration object created from the env.
type Config struct {
	Port     string
	LogPath  string
	LogLevel log.Level
}
