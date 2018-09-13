package db

import (
	"testing"
)

func TestCreateTag(t *testing.T) {
	holder := CreateTestDB()

	createTagIndex := "first-index"
	createTag := "first-tag"
	tag := holder.CreateTag(createTag, createTagIndex)
	if tag.Tagindex != createTagIndex {
		t.Error("expected: ", createTagIndex, " - got: ", tag.Tagindex)
	}
	if tag.Tag != createTag {
		t.Error("expected: ", createTag, " - got: ", tag.Tag)
	}
	if tag.ID != 1 {
		t.Error("expected id: 1 - got: ", tag.ID)
	}
}
