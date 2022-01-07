package article

import (
	"article/entity/request"
	"article/entity/response"
	"context"
)

type ArticleService interface {
	CreateArticle(ctx context.Context, article request.ArticleCreateRequest) response.ArticleCreateResponse
	GetArticles(ctx context.Context, params request.ArticleGetRequest) []response.ArticleGet
}
