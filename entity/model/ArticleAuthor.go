package model

import (
	"time"

	"github.com/google/uuid"
)

type ArticleAuthor struct {
	Id        uuid.UUID
	AuthorId  uuid.UUID
	Name      string
	Title     string
	Body      string
	CreatedAt time.Time
}
