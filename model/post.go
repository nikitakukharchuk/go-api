package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID    uuid.UUID `gorm:"type:uuid;"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
	BaseModel
}
type Posts struct {
	Posts []Post `json:"posts"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	post.ID = uuid.New()
	return
}
