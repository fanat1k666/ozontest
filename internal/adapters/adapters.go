package adapters

import "database/sql"

//go:generate mockgen -destination=mocks/db_mock.go . Db

type Db interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
}

type Storage interface {
	GetPost(page int, size int) []PostResult
	GetPostWithComments(postId int, page int, size int) []PostWithCommentsResult
	CreatePost(userId int, allowComments bool) error
	CreateComment(postId int, userId int, parId int, msg string) error
	Close() error
}

type PostWithCommentsResult struct {
	PostId            int
	PostAuthorId      int
	PostCreatedAt     string
	PostUpdatedAt     string
	PostAllowComments bool
	CommentId         int
	CommentAuthorId   int
	CommentParId      int
	CommentMsg        string
	CommentCreatedAt  string
	CommentUpdatedAt  string
}

type PostResult struct {
	Id            int
	AuthorId      int
	CreatedAt     string
	UpdatedAt     string
	AllowComments bool
}
