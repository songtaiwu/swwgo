package my_service

type ArticleAddRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
