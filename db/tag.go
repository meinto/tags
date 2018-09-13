package db

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Tagindex string
	Tag      string
}

func (db *DB) CreateTag(str string, index string) Tag {
	db.Instance.Create(&Tag{
		Tagindex: index,
		Tag:      str,
	})
	var tag Tag
	db.Instance.Last(&tag, "tag = ?", str)
	return tag
}

func (db *DB) CountTags(needle map[string]interface{}, index string) int {
	rows, err := db.Instance.Table("tags").Select("*").Where("tagindex = ?", index).Rows()
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return 0
	}

	count := 0
	for rows.Next() {
		found := true
		var tag Tag
		db.Instance.ScanRows(rows, &tag)

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
