package config

import (
	"github.com/felipear89/go-examples/pkg/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

type dsn string

func (s *dsn) replace(s1, s2 string) {
	*s = dsn(strings.ReplaceAll(s.string(), s1, s2))
}

func (s *dsn) string() string {
	return string(*s)
}

var (
	logLevel = map[string]logger.LogLevel{
		"silent": logger.Silent,
		"info":   logger.Info,
		"warn":   logger.Warn,
		"error":  logger.Error,
	}
)

func NewDatabase(config *Config) *gorm.DB {
	dsn := dsn("host={host} user={user} password='{password}' dbname={dbname} port={port} sslmode=disable TimeZone=UTC")
	dsn.replace("{host}", config.Database.Host)
	dsn.replace("{user}", config.Database.User)
	dsn.replace("{password}", config.Database.Password)
	dsn.replace("{dbname}", "go_examples")
	dsn.replace("{port}", "5432")

	db, err := gorm.Open(postgres.Open(dsn.string()), &gorm.Config{
		Logger: logger.New(log.NewEntry(log.StandardLogger()), logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel[config.Database.LogLevel],
		}),
	})

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.Database.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(config.Database.ConnMaxIdleTime)

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	err = db.AutoMigrate(&model.Publisher{}, &model.Book{})

	if err != nil {
		log.Fatal("Failed to run migration on database", err)
	}

	return db
}
