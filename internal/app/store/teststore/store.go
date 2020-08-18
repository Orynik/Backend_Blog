package teststore

import (
	"database/sql"

	"github.com/Kentabr5427/http-rest-api/internal/app/model"
	"github.com/Kentabr5427/http-rest-api/internal/app/store"
)

//Store ...
type Store struct {
	userRepository     *UserRepository
	db                 *sql.DB
	articleRepository  *ArticlesRepository
	commentRepository  *CommentRepository
	categoryRepository *CategoriesRepository
}

//New ...
func New() *Store {
	return &Store{}
}

//User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

//Article ...
func (s *Store) Article() store.ArticlesRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}

	s.articleRepository = &ArticlesRepository{
		store:    s,
		articles: make(map[int]*model.Article, 5),
	}

	return s.articleRepository
}

//Comment ...
func (s *Store) Comment() store.CommentRepository {
	if s.commentRepository != nil {
		return s.commentRepository
	}

	s.commentRepository = &CommentRepository{
		store:    s,
		comments: make(map[int]*model.Comment, 5),
	}

	return s.commentRepository
}

//Category ...
func (s *Store) Category() store.CategoryRepository {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}

	s.categoryRepository = &CategoriesRepository{
		store:      s,
		categories: make(map[int]*model.Category, 5),
	}

	return s.categoryRepository
}
