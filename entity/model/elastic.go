package model

import (
	"github.com/google/uuid"
)

type ElasticArticle struct {
	Id     uuid.UUID
	Author string
	Title  string
	Body   string
}
