package db

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Tag string
}

func (db *DB) CreateTag(str string) Tag {
	db.instance.Create(&Tag{
		Tag: str,
	})
	var tag Tag
	db.instance.Last(&tag, "tag = ?", str)
	return tag
}

func (db *DB) CountTags() int {
	// TODO
	return 0
}
