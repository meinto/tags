package db

import "github.com/jinzhu/gorm"

type DB struct {
	instance *gorm.DB
}

func New() *DB {
	instance, err := gorm.Open("sqlite3", "tags.db")
	if err != nil {
		panic("failed to connect database")
	}
	// defer db.Close()

	instance.AutoMigrate(&Tag{})

	return &DB{instance}
}
