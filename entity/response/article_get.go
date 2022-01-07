package response

import (
	"time"

	"github.com/google/uuid"
)

type ArticleGet struct {
	Id        uuid.UUID `json:"id"`
	AuthorId  uuid.UUID `json:"author_id"`
	Name      string    `validate:"required" json:"name"`
	Title     string    `validate:"required" json:"title"`
	Body      string    `validate:"required" json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
