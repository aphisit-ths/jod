package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbConfig() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
