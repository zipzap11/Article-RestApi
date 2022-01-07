package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ArticleController interface {
	GetArticles(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreateArticle(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
