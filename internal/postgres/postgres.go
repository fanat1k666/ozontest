package postgres

import (
	"database/sql"
	"fmt"
	"ozon/internal/adapters"

	_ "github.com/lib/pq"
)

type Db struct {
	adapters.Db
}

type Post struct {
	Id            int
	AuthorId      int
	CreatedAt     string
	UpdatedAt     sql.NullString
	AllowComments bool
}

type PostWithComments struct {
	PostId            int
	PostAuthorId      int
	PostCreatedAt     string
	PostUpdatedAt     sql.NullString
	PostAllowComments bool
	CommentId         int
	CommentAuthorId   int
	CommentParId      *int
	CommentMsg        string
	CommentCreatedAt  string
	CommentUpdatedAt  sql.NullString
}

func New(db adapters.Db) *Db {
	return &Db{db}
}

func ConnString(host string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)
}

func (db *Db) Close() error {
	return nil
}

func (pg *Db) GetPost(page int, size int) []adapters.PostResult {
	if page < 0 || size < 0 {
		return nil
	}
	offset := page * size
	rows, err := pg.Query(`SELECT * FROM "post"
		ORDER BY id
		LIMIT $1 OFFSET $2`, size, offset)
	if err != nil {
		fmt.Println("Cant request ", err)
	}

	var r Post
	var posts []Post
	for rows.Next() {
		err = rows.Scan(
			&r.Id,
			&r.AuthorId,
			&r.CreatedAt,
			&r.UpdatedAt,
			&r.AllowComments,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		posts = append(posts, r)
	}
	res := make([]adapters.PostResult, 0, len(posts))
	for i := range posts {
		res = append(res, adapters.PostResult{
			Id:            posts[i].Id,
			AuthorId:      posts[i].AuthorId,
			CreatedAt:     posts[i].CreatedAt,
			UpdatedAt:     posts[i].UpdatedAt.String,
			AllowComments: posts[i].AllowComments,
		})
	}

	return res
}

func (pg *Db) GetPostWithComments(postId int, page int, size int) []adapters.PostWithCommentsResult {
	if page < 0 || size < 0 {
		return nil
	}
	offset := page * size
	rows, err := pg.Query(`SELECT post.id,
       post.author_id,
       post.created_at,
       post.updated_at,
       post.allow_comments,
       comment.id,
       comment.author_id,
       comment.par_id,
       comment.msg,
       comment.created_at,
       comment.updated_at FROM "post"
        JOIN "comment"
            ON post.id = comment.post_id
            WHERE post.id = $1
            ORDER BY comment.created_at ASC 
            LIMIT $2 OFFSET $3`, postId, size, offset)
	if err != nil {
		fmt.Println("Cant request ", err)
	}

	var r PostWithComments
	var postsWithComments []PostWithComments

	for rows.Next() {
		var sqlCommentParId sql.NullInt32
		err = rows.Scan(
			&r.PostId,
			&r.PostAuthorId,
			&r.PostCreatedAt,
			&r.PostUpdatedAt,
			&r.PostAllowComments,
			&r.CommentId,
			&r.CommentAuthorId,
			&sqlCommentParId,
			&r.CommentMsg,
			&r.CommentCreatedAt,
			&r.CommentUpdatedAt,
		)
		if sqlCommentParId.Valid {
			t := int(sqlCommentParId.Int32)
			r.CommentParId = &t
		}
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}

		postsWithComments = append(postsWithComments, r)
	}
	var res = make([]adapters.PostWithCommentsResult, 0, len(postsWithComments))
	for i := range postsWithComments {
		var ParId int
		if postsWithComments[i].CommentParId != nil {
			ParId = *postsWithComments[i].CommentParId
		}
		res = append(res, adapters.PostWithCommentsResult{
			PostId:            postsWithComments[i].PostId,
			PostAuthorId:      postsWithComments[i].PostAuthorId,
			PostCreatedAt:     postsWithComments[i].PostCreatedAt,
			PostUpdatedAt:     postsWithComments[i].PostUpdatedAt.String,
			PostAllowComments: postsWithComments[i].PostAllowComments,
			CommentId:         postsWithComments[i].CommentId,
			CommentAuthorId:   postsWithComments[i].CommentAuthorId,
			CommentParId:      ParId,
			CommentMsg:        postsWithComments[i].CommentMsg,
			CommentCreatedAt:  postsWithComments[i].CommentCreatedAt,
			CommentUpdatedAt:  postsWithComments[i].CommentUpdatedAt.String,
		})
	}
	return res
}

func (pg *Db) CreatePost(userId int, allowComments bool) error {
	_, err := pg.Exec(`INSERT INTO post (author_id,allow_comments) VALUES ($1,$2)`, userId, allowComments)
	if err != nil {
		fmt.Println("Cant request ", err)
	}
	return nil
}

func (pg *Db) CreateComment(postId int, userId int, parId int, msg string) error {
	_, err := pg.Exec(`INSERT INTO comment (post_id,msg,par_id,author_id)
								SELECT id, $1,$2,$3 FROM post
								WHERE post.allow_comments = true AND
      							post.id = $4`, msg, parId, userId, postId)
	if err != nil {
		fmt.Println("Cant request ", err)
	}
	return nil
}
