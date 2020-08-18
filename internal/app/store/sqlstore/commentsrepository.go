package sqlstore

import (
	"github.com/Orynik/Backend_Blog/internal/app/model"
)

//CommentRepository ...
type CommentRepository struct {
	store *Store
}

//Create ...
func (s *CommentRepository) Create(a *model.Comment) error {
	v, err := s.store.db.Exec("INSERT INTO comments (`article`, `author`,`text`,`date`) VALUES(?,?,?,?)",
		&a.Article,
		&a.Author,
		&a.Text,
		&a.Date,
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

//GetComments  ...
func (s *CommentRepository) GetComments() ([]*model.Comment, error) {

	arrayOfComments := make([]*model.Comment, 0)

	row, err := s.store.db.Query("SELECT `id`, `article`, `author`, `text`, `date` FROM comments")

	if err != nil {
		return nil, err
	}

	for row.Next() {
		comment := &model.Comment{}
		err := row.Scan(&comment.ID, &comment.Article, &comment.Author, &comment.Text, &comment.Date)

		if err != nil {
			return nil, err
		}

		arrayOfComments = append(arrayOfComments, comment)
	}
	return arrayOfComments, nil
}
