package teststore

import (
	"github.com/Kentabr5427/http-rest-api/internal/app/model" // ...
)

//ArticlesRepository ...
type ArticlesRepository struct {
	store    *Store
	articles map[int]*model.Article
}

//Create ...
func (r *ArticlesRepository) Create(a *model.Article) error {
	a.ID = len(r.articles) + 1
	r.articles[a.ID] = a
	return nil
}

//GetArticles ...
func (r *ArticlesRepository) GetArticles() ([]*model.Article, error) {

	array := make([]*model.Article, len(r.articles)+1)
	for idx, item := range r.articles {
		array[idx-1] = item
	}

	return array, nil
}
