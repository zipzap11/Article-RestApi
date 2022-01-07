package helper

import (
	"article/entity/model"
	"article/entity/response"
)

func ToArticleCreateResponse(article model.Article) response.ArticleCreateResponse {
	return response.ArticleCreateResponse{
		Id:        article.Id,
		AuthorId:  article.AuthorId,
		Title:     article.Title,
		Body:      article.Body,
		CreatedAt: article.CreatedAt,
	}
}
