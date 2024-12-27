package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}
type Posts struct {
	Posts []Post `json:"posts"`
}

//func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
//	post.ID = uuid.New()
//	return
//}
