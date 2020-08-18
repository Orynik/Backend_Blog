package sqlstore

import (
	"github.com/Orynik/Backend_Blog/internal/app/model"
)

//ArticlesRepository ...
type ArticlesRepository struct {
	store *Store
}

//Create ...
func (r *ArticlesRepository) Create(a *model.Article) error {

	v, err := r.store.db.Exec("INSERT INTO articles (`title`, `date`,`category`,`author`,`content`) VALUES(?,?,?,?,?)",
		&a.Title,
		&a.Date,
		&a.Category,
		&a.Author,
		&a.Content,
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

//GetArticles ...
func (r *ArticlesRepository) GetArticles() ([]*model.Article, error) {

	arrayOfArticles := make([]*model.Article, 0)

	row, err := r.store.db.Query("SELECT `id`, `title`, `date`, `category`, `author`, `content` FROM articles")

	if err != nil {
		return nil, err
	}

	for row.Next() {
		article := &model.Article{}
		err := row.Scan(&article.ID, &article.Title, &article.Date, &article.Category, &article.Author, &article.Content)

		if err != nil {
			return nil, err
		}

		arrayOfArticles = append(arrayOfArticles, article)
	}
	return arrayOfArticles, nil
}
