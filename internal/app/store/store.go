package store

//Store ..
type Store interface {
	User() UserRepository
	Article() ArticlesRepository
	Comment() CommentRepository
	Category() CategoryRepository
}
