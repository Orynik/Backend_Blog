package model

//Comment ...
type Comment struct {
	ID      int    `json:"id"`
	Article int    `json:"idArticle"`
	Author  string `json:"author"`
	Text    string `json:"text"`
	Date    string `json:"date"`
}
