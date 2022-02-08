package app

import "github.com/google/uuid"

type Article struct {
	Uuid    uuid.UUID `json:"uuid"`
	Title   string
	Content string
}

func (Article) TableName() string {
	return "articles"
}
