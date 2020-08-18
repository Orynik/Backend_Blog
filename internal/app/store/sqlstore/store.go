package sqlstore

import (
	"database/sql"

	"github.com/Orynik/Backend_Blog/internal/app/store"
	_ "github.com/go-sql-driver/mysql" // ...
)

//Store ...
type Store struct {
	db                 *sql.DB
	userRepository     *UserRepository
	articleRepository  *ArticlesRepository
	commentRepository  *CommentRepository
	categoryRepository *CategoryRepository
}

//New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

//Article ...
func (s *Store) Article() store.ArticlesRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}

	s.articleRepository = &ArticlesRepository{
		store: s,
	}

	return s.articleRepository
}

//Comment ...
func (s *Store) Comment() store.CommentRepository {
	if s.commentRepository != nil {
		return s.commentRepository
	}

	s.commentRepository = &CommentRepository{
		store: s,
	}

	return s.commentRepository
}

//Category ...
func (s *Store) Category() store.CategoryRepository {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}

	s.categoryRepository = &CategoryRepository{
		store: s,
	}

	return s.categoryRepository
}
