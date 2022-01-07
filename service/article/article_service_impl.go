package article

import (
	"article/elasticsearch"
	"article/entity/model"
	"article/entity/request"
	"article/entity/response"
	"article/exception"
	"article/helper"
	article_repo "article/repository/article"
	author_repo "article/repository/author"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ArticleServiceImpl struct {
	DB               *sql.DB
	ArticleRepo      article_repo.ArticleRepository
	AuthorRepo       author_repo.AuthorRepository
	Validate         *validator.Validate
	SearchRepository elasticsearch.SearchRepository
}

func NewArticleService(DB *sql.DB, articleRepo article_repo.ArticleRepository, authorRepo author_repo.AuthorRepository, validate *validator.Validate, searchRepo elasticsearch.SearchRepository) ArticleService {
	return &ArticleServiceImpl{
		DB:               DB,
		ArticleRepo:      articleRepo,
		AuthorRepo:       authorRepo,
		Validate:         validate,
		SearchRepository: searchRepo,
	}
}

func (service *ArticleServiceImpl) CreateArticle(ctx context.Context, request request.ArticleCreateRequest) response.ArticleCreateResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommmitOrRollback(tx)

	authorData := helper.ToAuthorModel(request)

	if authorData.Id == uuid.Nil {
		authorData.Id = uuid.New()

		_, err := service.AuthorRepo.Create(ctx, tx, authorData)
		helper.PanicIfErr(err)

	} else {
		fmt.Println("author id = ", authorData.Id)
		authorById, err := service.AuthorRepo.GetById(ctx, tx, authorData.Id)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		if authorById.Id == uuid.Nil {
			panic(exception.NewNotFoundError("not found"))
		}
	}
	request.AuthorId = authorData.Id.String()

	articleData := helper.ToArticleModel(request)
	articleData.Id = uuid.New()
	articleData.CreatedAt = time.Now()

	createdData := service.ArticleRepo.Create(ctx, tx, articleData)

	elasticData := model.ElasticArticle{
		Id:     createdData.Id,
		Author: request.Name,
		Title:  createdData.Title,
		Body:   createdData.Body,
	}
	service.SearchRepository.Insert(ctx, elasticData)

	return helper.ToArticleCreateResponse(createdData)
}

func (service *ArticleServiceImpl) GetArticles(ctx context.Context, request request.ArticleGetRequest) []response.ArticleGet {
	fmt.Println("request = ", request)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommmitOrRollback(tx)

	if request.Author == "" && request.Query == "" {
		fmt.Println("masuk if")
		articles := service.ArticleRepo.GetAll(ctx, tx)
		fmt.Println("articles ==", articles)
		return helper.ToArticleResponseSlice(articles)
	}
	fmt.Println("lolos if")
	result := service.SearchRepository.Query(ctx, request)
	var ids []uuid.UUID
	fmt.Println("result from ES", result)
	for _, data := range result {
		ids = append(ids, data.Id)
	}

	articles := service.ArticleRepo.GetByMultiID(ctx, tx, ids)
	helper.PanicIfErr(err)

	return helper.ToArticleResponseSlice(articles)
}
