package elasticsearch

import (
	"article/entity/model"
	"article/entity/request"
	"article/helper"
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type ElasticSearchRepositoryImpl struct {
}

func NewSearchRepository() SearchRepository {
	return &ElasticSearchRepositoryImpl{}
}

func (repository *ElasticSearchRepositoryImpl) getClient() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false), elastic.SetHealthcheck(false))
	helper.PanicIfErr(err)

	fmt.Println("Success Initialized ES")

	return client
}

func (repository *ElasticSearchRepositoryImpl) Insert(ctx context.Context, article model.ElasticArticle) {
	fmt.Println("insert articles elastic", article)
	esclient := repository.getClient()

	jsonData, err := json.Marshal(article)
	helper.PanicIfErr(err)

	json := string(jsonData)

	_, err = esclient.Index().Index("articles").BodyJson(json).Do(ctx)
	helper.PanicIfErr(err)
}

func (repository *ElasticSearchRepositoryImpl) Query(ctx context.Context, param request.ArticleGetRequest) []model.ElasticArticle {
	fmt.Println("query elastic", param)
	esclient := repository.getClient()

	var articles []model.ElasticArticle

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("title", param.Query))
	searchSource.Query(elastic.NewMatchQuery("body", param.Query))
	searchSource.Query(elastic.NewMatchQuery("author", param.Author))

	queryStr, err := searchSource.Source()
	helper.PanicIfErr(err)

	queryJson, err := json.Marshal(queryStr)
	fmt.Println("queryJson === ", string(queryJson))
	helper.PanicIfErr(err)

	searchService := esclient.Search().Index("articles").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	helper.PanicIfErr(err)
	fmt.Println("searchresult = ", searchResult)
	for _, hit := range searchResult.Hits.Hits {
		var article model.ElasticArticle
		err := json.Unmarshal(hit.Source, &article)
		helper.PanicIfErr(err)

		articles = append(articles, article)
	}

	return articles
}
