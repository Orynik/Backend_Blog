package teststore_test

import (
	"testing"

	"github.com/Kentabr5427/http-rest-api/internal/app/model"

	"github.com/Kentabr5427/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestArticlesRepository_Create(t *testing.T) {
	s := teststore.New()
	a := model.TestArticles(t)
	err := s.Article().Create(a)
	assert.NoError(t, err)
}

func TestArticlesRepository_GetArticles(t *testing.T) {
	//Запись первого пользователя
	s := teststore.New()
	a := model.TestArticles(t)
	_ = s.Article().Create(a)

	//Запись второго пользователя
	a = model.TestArticles(t)
	a.Author = "Orynik"
	a.Title = "Vue.js"
	_ = s.Article().Create(a)

	result, err := s.Article().GetArticles()
	assert.NoError(t, err)
	//Проверка корректности первой записи
	assert.Equal(t, result[0].Title, "Golang")
	assert.Equal(t, result[0].Content, "LoremIpsum")
	assert.Equal(t, result[0].Category, "Backend")

	//Проверка корректности второй записи
	assert.Equal(t, result[1].Title, "Vue.js")
	assert.Equal(t, result[1].Author, "Orynik")
	assert.Equal(t, result[1].Category, "Backend")

	//{2 Vue.js 2020-05-28T01:32:29.413Z Backend Orynik LoremIpsum}
	//{1 Golang 2020-05-28T01:32:29.413Z Backend Cookies@chort.ru LoremIpsum}
}
