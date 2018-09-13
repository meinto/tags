package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func CreateTestDB() *DB {
	err := os.Remove("../test.db")
	if err != nil {
		log.Println("error while deleting test.db: ", err)
	}
	instance, err := gorm.Open("sqlite3", "../test.db")
	if err != nil {
		log.Panicln(err)
		panic("failed to connect database")
	}
	// defer db.Close()

	instance.AutoMigrate(&Tag{})

	return &DB{instance}
}

func FillWithMockData(holder *DB) {
	holder.Instance.Create(&Tag{Tagindex: "mock-index-1", Tag: "mock-tag-1"})
	holder.Instance.Create(&Tag{Tagindex: "mock-index-1", Tag: "mock-tag-1"})
	holder.Instance.Create(&Tag{Tagindex: "mock-index-1", Tag: "mock-tag-2"})
	holder.Instance.Create(&Tag{Tagindex: "mock-index-1", Tag: "mock-tag-2"})

	holder.Instance.Create(&Tag{Tagindex: "mock-index-2", Tag: "mock-tag-2"})

	holder.Instance.Create(&Tag{Tagindex: "mock-index-3", Tag: "mock-tag-2"})
	holder.Instance.Create(&Tag{Tagindex: "mock-index-3", Tag: "mock-tag-3"})
	holder.Instance.Create(&Tag{Tagindex: "mock-index-3", Tag: "mock-tag-3"})
	holder.Instance.Create(&Tag{Tagindex: "mock-index-3", Tag: "mock-tag-3"})
	holder.Instance.Create(&Tag{Tagindex: "mock-index-3", Tag: "mock-tag-3"})
}
