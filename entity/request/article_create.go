package request

type ArticleCreateRequest struct {
	AuthorId string `json:"author_id"`
	Name     string `validate:"required" json:"name"`
	Title    string `validate:"required" json:"title"`
	Body     string `validate:"required" json:"body"`
}
