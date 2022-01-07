package repository

import (
	"article/entity/model"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type ArticleRepository interface {
	Create(ctx context.Context, tx *sql.Tx, article model.Article) model.Article
	GetByMultiID(ctx context.Context, tx *sql.Tx, id []uuid.UUID) []model.ArticleAuthor
	GetAll(ctx context.Context, tx *sql.Tx) []model.ArticleAuthor
}
