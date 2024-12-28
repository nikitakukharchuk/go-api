package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Posts struct {
	Posts []Post `json:"posts"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	post.ID = uuid.New()
	return
}
