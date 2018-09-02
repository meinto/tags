package db

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
}

func (db *DB) CreateTag() *Tag {
	// TODO
	return &Tag{}
}

func (db *DB) CountTags() int {
	// TODO
	return 0
}
