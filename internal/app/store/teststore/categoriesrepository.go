package teststore

import "github.com/Kentabr5427/http-rest-api/internal/app/model"

//CategoriesRepository ...
type CategoriesRepository struct {
	store      *Store
	categories map[int]*model.Category
}

//Create ...
func (r *CategoriesRepository) Create(a *model.Category) error {
	a.ID = len(r.categories) + 1
	r.categories[a.ID] = a
	return nil
}

//GetCategories ...
func (r *CategoriesRepository) GetCategories() ([]*model.Category, error) {

	array := make([]*model.Category, len(r.categories)+1)
	for idx, item := range r.categories {
		array[idx-1] = item
	}

	return array, nil
}
