package model

import "testing"

//TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

//TestArticles ...
func TestArticles(t *testing.T) *Article {
	return &Article{
		Title:    "Golang",
		Date:     "2020-05-28T01:32:29.413Z",
		Category: "Backend",
		Author:   "Cookies@chort.ru",
		Content:  "LoremIpsum",
	}
}

//TestComment ...
func TestComment(t *testing.T) *Comment {
	return &Comment{
		ID:      0,
		Article: 123,
		Author:  "qwe",
		Text:    "qwe",
		Date:    "qwe",
	}
}

//TestCategory ...
func TestCategory(t *testing.T) *Category {
	return &Category{
		ID:       0,
		Category: "Backend",
	}
}
