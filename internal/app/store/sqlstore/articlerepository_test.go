package sqlstore_test

import (
	"testing"

	"github.com/Orynik/Backend_Blog/internal/app/model"
	"github.com/Orynik/Backend_Blog/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestArticleRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("articles")

	s := sqlstore.New(db)
	a := model.TestArticles(t)

	assert.NoError(t, s.Article().Create(a))
	assert.NotNil(t, a)
	assert.Equal(t, a.ID, 1)
}

func TestArticleRepository_GetArticles(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("articles")

	s := sqlstore.New(db)

	//Создание и запись первой статьи

	a := model.TestArticles(t)
	_ = s.Article().Create(a)

	//Создание и запись второй статьи
	a = model.TestArticles(t)
	a.Title = "Test"
	a.Author = "Orynik"
	_ = s.Article().Create(a)

	result, err := s.Article().GetArticles()

	assert.NoError(t, err)
	assert.Equal(t, len(result), 2)

	//Проверка значений
	assert.Equal(t, result[1].Title, "Test")
	assert.Equal(t, result[1].Author, "Orynik")

}
