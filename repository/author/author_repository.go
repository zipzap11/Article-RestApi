package author

import (
	"article/entity/model"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type AuthorRepository interface {
	Create(ctx context.Context, tx *sql.Tx, author model.Author) (model.Author, error)
	GetById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (model.Author, error)
}
