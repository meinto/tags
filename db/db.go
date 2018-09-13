package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	Instance *gorm.DB
}

func New() *DB {
	instance, err := gorm.Open("sqlite3", "tags.db")
	if err != nil {
		log.Panicln(err)
		panic("failed to connect database")
	}
	// defer db.Close()

	instance.AutoMigrate(&Tag{})

	return &DB{instance}
}
