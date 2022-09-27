package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	Pg *gorm.DB
}

func GetDB(connectionString string) *DB {
	open, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return &DB{
		Pg: open,
	}
}

func (d DB) Migrate(model ...any) {
	err := d.Pg.AutoMigrate(model...)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
