package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"selfletter-backend/config"
	"time"
)

var databaseHandle *gorm.DB

type User struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

type Topic struct {
	Name string `json:"name"`
}

type AdminKey struct {
	Key string `json:"key"`
}

type UsersTopicsRel struct {
	Email string `json:"email"`
	Topic string `json:"topic"`
}

func Open() error {
	cfg := config.GetConfig()
	if cfg.Debug == false {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  false,
			},
		)

		open, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  cfg.DSN,
			PreferSimpleProtocol: true,
		}), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			return err
		}

		databaseHandle = open
	} else {
		open, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  cfg.DSN,
			PreferSimpleProtocol: true,
		}))

		if err != nil {
			return err
		}

		databaseHandle = open
	}

	err := databaseHandle.AutoMigrate(&User{}, &UsersTopicsRel{}, &Topic{}, &AdminKey{})
	if err != nil {
		return err
	}

	return nil
}

func GetDatabaseHandle() *gorm.DB {
	return databaseHandle
}
