package database

import "gorm.io/gorm"

const (
	CONDITION_FIELD_POST_TITLE     = "Title"
	CONDITION_FIELD_POST_UPDATEDAT = "UpdatedAt"
	CONDITION_FIELD_POST_AUTHOR    = "Author"
)

type ModelPost struct {
	gorm.Model
	UUID         string `gorm:"uniqueIndex"`
	Title        string
	Context      string
	Author       string
	PasswordHash string
}
