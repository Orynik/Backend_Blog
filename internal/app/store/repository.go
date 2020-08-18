package store

import "github.com/Orynik/Backend_Blog/internal/app/model"

//UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

//ArticlesRepository ...
type ArticlesRepository interface {
	Create(*model.Article) error
	GetArticles() ([]*model.Article, error)
}

//CommentRepository ...
type CommentRepository interface {
	Create(*model.Comment) error
	GetComments() ([]*model.Comment, error)
}

//CategoryRepository ...
type CategoryRepository interface {
	Create(*model.Category) error
	GetCategories() ([]*model.Category, error)
}
