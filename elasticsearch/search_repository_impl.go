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
	Host string
	Port string
}

func NewSearchRepository(host string, port string) SearchRepository {
	return &ElasticSearchRepositoryImpl{
		Host: host,
		Port: port,
	}
}

func (repository *ElasticSearchRepositoryImpl) getClient() *elastic.Client {
	URL := fmt.Sprintf("http://%v:%v", repository.Host, repository.Port)
	client, err := elastic.NewClient(elastic.SetURL(URL), elastic.SetSniff(false), elastic.SetHealthcheck(false))
	helper.PanicIfErr(err)

	fmt.Println("Success Initialized ES")

	return client
}

func (repository *ElasticSearchRepositoryImpl) Insert(ctx context.Context, article model.ElasticArticle) {
	esclient := repository.getClient()

	jsonData, err := json.Marshal(article)
	helper.PanicIfErr(err)

	json := string(jsonData)

	_, err = esclient.Index().Index("articles").BodyJson(json).Id(article.Id.String()).Do(ctx)
	helper.PanicIfErr(err)
}

func (repository *ElasticSearchRepositoryImpl) Query(ctx context.Context, param request.ArticleGetRequest) []model.ElasticArticle {
	esclient := repository.getClient()

	var articles []model.ElasticArticle

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMultiMatchQuery(param.Query, "Title", "Body"))
	searchSource.Query(elastic.NewBoolQuery().Should(
		elastic.NewMultiMatchQuery(param.Query, "Title", "Body"),
		elastic.NewMatchQuery("Author", param.Query),
		elastic.NewWildcardQuery("Author", param.Author),
		elastic.NewWildcardQuery("Title", param.Query),
		elastic.NewWildcardQuery("Body", param.Query),
	))

	// queryStr, err := searchSource.Source()
	// helper.PanicIfErr(err)

	// queryJson, err := json.Marshal(queryStr)
	// helper.PanicIfErr(err)

	searchService := esclient.Search().Index("articles").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	helper.PanicIfErr(err)
	for _, hit := range searchResult.Hits.Hits {
		var article model.ElasticArticle
		err := json.Unmarshal(hit.Source, &article)
		helper.PanicIfErr(err)

		articles = append(articles, article)
	}

	return articles
}
