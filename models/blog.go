package models

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	Id         string     `bson:"id" json:"id"`
	Title      string     `bson:"title" json:"title"`
	Content    string     `bson:"content" json:"content"`
	Category   string     `bson:"category,omitempty" json:"category"`
	Tags       []string   `bson:"tags,omitempty" tags:"tags"`
	CreatedAt  *time.Time `bson:"createdAt" json:"createdAt"`
	ModifiedAt *time.Time `bson:"modifiedAt" json:"modifiedAt"`
}

func New(title, content, category string, tags []string) Blog {
	now := time.Now()
	return Blog{
		Id:         uuid.NewString(),
		Title:      title,
		Content:    content,
		Category:   category,
		Tags:       tags,
		CreatedAt:  &now,
		ModifiedAt: &now,
	}
}
