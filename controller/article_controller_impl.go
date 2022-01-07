package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"article/cache"
	req "article/entity/request"
	"article/entity/response"
	"article/helper"
	article_service "article/service/article"
)

type ArticleControllerImpl struct {
	ArticleService article_service.ArticleService
	ArticleCache   cache.Cache
}

func NewArticleService(articleService article_service.ArticleService, articleCache cache.Cache) ArticleController {
	return &ArticleControllerImpl{
		ArticleService: articleService,
		ArticleCache:   articleCache,
	}
}

func (controller *ArticleControllerImpl) GetArticles(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	queryParams := request.URL.Query()
	author := queryParams.Get("author")
	query := queryParams.Get("query")

	key := query + ":" + author

	var responseData []response.ArticleGet

	// lookup data from cache first
	responseData = controller.ArticleCache.Get(key)

	if len(responseData) == 0 {
		articleRequestData := req.ArticleGetRequest{
			Query:  query,
			Author: author,
		}

		data := controller.ArticleService.GetArticles(request.Context(), articleRequestData)

		controller.ArticleCache.Set(key, data)
		responseData = data
	}

	response := response.StandardResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   responseData,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	helper.PanicIfErr(err)
}

func (controller *ArticleControllerImpl) CreateArticle(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	articleRequestData := req.ArticleCreateRequest{}
	err := decoder.Decode(&articleRequestData)
	helper.PanicIfErr(err)

	result := controller.ArticleService.CreateArticle(request.Context(), articleRequestData)

	response := response.StandardResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(response)
	helper.PanicIfErr(err)
}
