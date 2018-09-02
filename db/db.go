package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	instance *gorm.DB
}

func New() *DB {
	f, _ := os.OpenFile("tags.db", os.O_RDONLY|os.O_CREATE, 0775)
	defer f.Close()
	instance, err := gorm.Open("sqlite3", "tags.db")
	if err != nil {
		log.Panicln(err)
		panic("failed to connect database")
	}
	// defer db.Close()

	instance.AutoMigrate(&Tag{})

	return &DB{instance}
}
