package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func GetDB(connectionString string) *gorm.DB {
	if db == nil {
		open, err := gorm.Open(postgres.Open(connectionString))
		if err != nil {
			log.Fatal(err.Error())
			return nil
		}

		db = open

		GetDB(connectionString)
	}

	return db
}
