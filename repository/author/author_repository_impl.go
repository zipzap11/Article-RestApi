package author

import (
	"article/entity/model"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type AuthorRepositoryImpl struct{}

func NewAuthorRepository() AuthorRepository {
	return &AuthorRepositoryImpl{}
}

func (repository *AuthorRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, author model.Author) (model.Author, error) {
	SQL := "INSERT INTO authors (id, name) VALUES ($1, $2)"

	_, err := tx.ExecContext(ctx, SQL, author.Id, author.Name)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (repository *AuthorRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (model.Author, error) {
	SQL := "SELECT id, name FROM authors WHERE id = $1 LIMIT 1"

	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		return model.Author{}, err
	}
	defer rows.Close()

	var authorData model.Author
	if rows.Next() {
		rows.Scan(authorData.Id, authorData.Name)
	}

	return authorData, nil
}
