package model

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	Id        uuid.UUID
	AuthorId  uuid.UUID
	Title     string
	Body      string
	CreatedAt time.Time
}
