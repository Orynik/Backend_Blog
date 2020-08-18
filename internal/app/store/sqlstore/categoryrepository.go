package sqlstore

import "github.com/Orynik/Backend_Blog/internal/app/model"

//CategoryRepository ...
type CategoryRepository struct {
	store *Store
}

//Create ...
func (r *CategoryRepository) Create(a *model.Category) error {

	v, err := r.store.db.Exec("INSERT INTO category (`category`) VALUES(?)",
		&a.Category,
	)

	if err != nil {
		return err
	}

	resultID, err := v.LastInsertId()

	a.ID = int(resultID)

	if err != nil {
		return nil
	}

	return nil
}

//GetCategories ...
func (r *CategoryRepository) GetCategories() ([]*model.Category, error) {

	arrayOfCategory := make([]*model.Category, 0)

	row, err := r.store.db.Query("SELECT `id`, `category` FROM category")

	if err != nil {
		return nil, err
	}

	for row.Next() {
		category := &model.Category{}
		err := row.Scan(&category.ID, &category.Category)

		if err != nil {
			return nil, err
		}

		arrayOfCategory = append(arrayOfCategory, category)
	}
	return arrayOfCategory, nil
}
