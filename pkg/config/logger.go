package config

import (
	log "github.com/sirupsen/logrus"
)

var (
	level = map[string]log.Level{
		"debug": log.DebugLevel,
		"info":  log.InfoLevel,
		"warn":  log.WarnLevel,
		"error": log.ErrorLevel,
		"fatal": log.FatalLevel,
		"panic": log.PanicLevel,
	}
	formatter = map[string]log.Formatter{
		"text": &log.TextFormatter{},
		"json": &log.JSONFormatter{},
	}
)

func SetupLogrus(config *Config) *log.Logger {
	log.SetLevel(level[config.LogLevel])
	log.SetFormatter(formatter[config.LogFormatter])
	return log.StandardLogger()
}
