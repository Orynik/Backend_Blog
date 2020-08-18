package model

//Article ...
type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Category string `json:"category"`
	Author   string `json:"author"`
	Content  string `json:"content"`
}

/*TODO: Прописать дальнейшие методы, которые нужно будет выполнить с Article
Допустим, валидацию даты*/
