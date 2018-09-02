package db

import (
	"encoding/json"
	"strings"

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

func (db *DB) CountTags(needle map[string]interface{}) int {
	rows, err := db.instance.Table("tags").Select("*").Rows()
	defer rows.Close()
	if err != nil {
		return -1
	}

	count := 0
	for rows.Next() {
		found := true
		var tag Tag
		db.instance.ScanRows(rows, &tag)

		var haystack map[string]interface{}
		json.NewDecoder(strings.NewReader(tag.Tag)).Decode(&haystack)

		for k := range needle {
			if needle[k] != haystack[k] {
				found = false
			}
		}

		if found {
			count++
		}
	}
	return count
}
