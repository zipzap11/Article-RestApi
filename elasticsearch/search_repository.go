package elasticsearch

import (
	"article/entity/model"
	"article/entity/request"
	"context"
)

type SearchRepository interface {
	Insert(ctx context.Context, article model.ElasticArticle)
	Query(ctx context.Context, param request.ArticleGetRequest) []model.ElasticArticle
}
