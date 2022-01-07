package request

type ArticleGetRequest struct {
	Query  string `json:"query"`
	Author string `json:"author"`
}
