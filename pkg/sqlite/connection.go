package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("quran.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
