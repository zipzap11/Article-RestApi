package response

import (
	"time"

	"github.com/google/uuid"
)

type ArticleCreateResponse struct {
	Id        uuid.UUID `json:"id"`
	AuthorId  uuid.UUID `json:"author_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"create_at"`
}
