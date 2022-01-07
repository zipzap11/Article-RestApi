package helper

import (
	"article/entity/model"
	"article/entity/request"
	"article/entity/response"

	"github.com/google/uuid"
)

func ToArticleModel(data request.ArticleCreateRequest) model.Article {
	var id uuid.UUID
	if len(data.AuthorId) == 0 {
		id = uuid.Nil
	} else {
		id = uuid.MustParse(data.AuthorId)
	}
	return model.Article{
		AuthorId: id,
		Title:    data.Title,
		Body:     data.Body,
	}
}

func ToAuthorModel(data request.ArticleCreateRequest) model.Author {
	var id uuid.UUID
	if len(data.AuthorId) == 0 {
		id = uuid.Nil
	} else {
		id = uuid.MustParse(data.AuthorId)
	}
	return model.Author{
		Id:   id,
		Name: data.Name,
	}
}

func ToArticleResponse(data model.ArticleAuthor) response.ArticleGet {
	return response.ArticleGet{
		Id:        data.Id,
		AuthorId:  data.AuthorId,
		Name:      data.Name,
		Title:     data.Title,
		Body:      data.Body,
		CreatedAt: data.CreatedAt,
	}
}

func ToArticleResponseSlice(data []model.ArticleAuthor) []response.ArticleGet {
	converted := []response.ArticleGet{}
	for _, val := range data {
		converted = append(converted, ToArticleResponse(val))
	}
	return converted
}
