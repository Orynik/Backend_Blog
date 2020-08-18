package teststore

import "github.com/Orynik/Backend_Blog/internal/app/model"

//CommentRepository ...
type CommentRepository struct {
	store    *Store
	comments map[int]*model.Comment
}

//Create ...
func (r *CommentRepository) Create(a *model.Comment) error {
	a.ID = len(r.comments) + 1
	r.comments[a.ID] = a
	return nil
}

//GetComments ...
func (r *CommentRepository) GetComments() ([]*model.Comment, error) {

	array := make([]*model.Comment, len(r.comments)+1)
	for idx, item := range r.comments {
		array[idx-1] = item
	}

	return array, nil
}
