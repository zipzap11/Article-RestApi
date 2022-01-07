package main

import (
	"article/cache"
	article_controller "article/controller"
	"article/database"
	"article/elasticsearch"
	"article/exception"
	"article/helper"
	article_repository "article/repository/article"
	author_repository "article/repository/author"
	article_service "article/service/article"
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// OR: Read migrations from a folder:
	DB := database.InitDB()
	ctx := context.Background()

	// repository
	articleRepository := article_repository.NewArticleRepository()
	authorRepository := author_repository.NewAuthorRepository()
	searchRepository := elasticsearch.NewSearchRepository()

	// service
	articleService := article_service.NewArticleService(DB, articleRepository, authorRepository, validator.New(), searchRepository)

	// cache
	articleRedisCache := cache.NewRedisCacheImpl(ctx, "localhost:6379", 0, 1*time.Hour)

	// controller
	articleController := article_controller.NewArticleService(articleService, articleRedisCache)

	router := httprouter.New()

	// router
	router.GET("/articles", articleController.GetArticles)
	router.POST("/articles", articleController.CreateArticle)

	// handle if panic happen
	router.PanicHandler = exception.ErrorHandler

	// server
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
