package cache

import (
	"article/entity/response"
)

type Cache interface {
	Set(key string, value []response.ArticleGet)
	Get(key string) []response.ArticleGet
}
