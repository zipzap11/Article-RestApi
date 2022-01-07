package repository

import (
	"context"
	"database/sql"
	"fmt"

	"article/entity/model"
	"article/exception"
	"article/helper"

	"github.com/google/uuid"
)

type ArticleRepositoryImpl struct{}

func NewArticleRepository() ArticleRepository {
	return &ArticleRepositoryImpl{}
}

func (repository *ArticleRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, article model.Article) model.Article {
	SQL := "INSERT INTO articles (id, author_id, title, body, created_at) VALUES($1, $2, $3, $4, $5)"

	_, err := tx.ExecContext(ctx, SQL, article.Id, article.AuthorId, article.Title, article.Body, article.CreatedAt)
	helper.PanicIfErr(err)

	return article
}

func (repository *ArticleRepositoryImpl) GetByMultiID(ctx context.Context, tx *sql.Tx, ids []uuid.UUID) []model.ArticleAuthor {
	if len(ids) == 0 {
		panic(exception.NewNotFoundError("not found author id"))
	}

	fmt.Println("inject", helper.DynamicInject(len(ids)))

	SQL := fmt.Sprintf(`SELECt articles.id, author_id, title, body, created_at, authors.name FROM articles 
			INNER JOIN authors ON articles.author_id = authors.id WHERE articles.id IN 
			%v ORDER BY articles.created_at DESC LIMIT 50;`, helper.DynamicInject(len(ids)))

	rows, err := tx.QueryContext(ctx, SQL, ids)
	helper.PanicIfErr(err)
	defer rows.Close()

	var articles []model.ArticleAuthor
	for rows.Next() {
		temp := model.ArticleAuthor{}
		rows.Scan(temp.Id, temp.AuthorId, temp.Title, temp.Body, temp.CreatedAt, temp.Name)

		articles = append(articles, temp)
	}

	return articles
}

func (repository *ArticleRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []model.ArticleAuthor {
	SQL := `SELECT articles.id, authors.id, title, body, created_at, authors.name FROM articles
			INNER JOIN authors ON articles.author_id = authors.id ORDER BY articles.created_at DESC LIMIT 50;`

	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	var articles []model.ArticleAuthor
	for rows.Next() {
		temp := model.ArticleAuthor{}
		rows.Scan(&temp.Id, &temp.AuthorId, &temp.Title, &temp.Body, &temp.CreatedAt, &temp.Name)
		articles = append(articles, temp)
	}

	return articles
}
