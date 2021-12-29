package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Environment  string
	Port         string
	Database     Database
	LogLevel     string
	LogFormatter string
}

type Database struct {
	Host            string
	Port            string
	User            string
	Password        string
	Name            string
	LogLevel        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

func init() {
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
}

func NewConfig() *Config {
	c := &Config{
		Environment:  viper.GetString("ENV"),
		Port:         viper.GetString("PORT"),
		LogLevel:     viper.GetString("LOG_LEVEL"),
		LogFormatter: viper.GetString("LOG_FORMATTER"),
		Database: Database{
			Host:            viper.GetString("DB_HOST"),
			Port:            viper.GetString("DB_PORT"),
			User:            viper.GetString("DB_USER"),
			Password:        viper.GetString("DB_PASSWORD"),
			Name:            viper.GetString("DB_NAME"),
			LogLevel:        viper.GetString("DB_LOG_LEVEL"),
			MaxIdleConns:    viper.GetInt("DB_MAX_IDLE_CONNS"),
			MaxOpenConns:    viper.GetInt("DB_MAX_OPEN_CONNS"),
			ConnMaxLifetime: viper.GetDuration("DB_CONN_MAX_LIFETIME"),
			ConnMaxIdleTime: viper.GetDuration("DB_CONN_MAX_IDLE_TIME"),
		},
	}
	SetupLogrus(c)
	return c
}
