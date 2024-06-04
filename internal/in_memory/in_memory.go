package in_memory

import (
	"fmt"
	"ozon/internal/adapters"
	"sort"
	"time"
)

type Comment struct {
	Id        int
	PostId    int
	Msg       string
	ParId     *int
	AuthorId  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func pointerTo[T any](in T) *T {
	return &in
}

type Im struct {
	posts   map[int]adapters.PostResult
	comment map[int]Comment
}

func (im *Im) Close() error {
	return nil
}

func New() (*Im, error) {
	return &Im{
		posts: map[int]adapters.PostResult{
			1: {
				Id:            1,
				AuthorId:      1,
				CreatedAt:     time.Now().String(),
				UpdatedAt:     "",
				AllowComments: false,
			},
			2: {
				Id:            2,
				AuthorId:      2,
				CreatedAt:     time.Now().Add(time.Second).String(),
				UpdatedAt:     "",
				AllowComments: true,
			},
		},
		comment: map[int]Comment{
			1: {
				Id:        1,
				PostId:    2,
				Msg:       "1",
				ParId:     nil,
				AuthorId:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Time{},
			},
			2: {
				Id:        2,
				PostId:    2,
				Msg:       "2",
				ParId:     pointerTo(1),
				AuthorId:  2,
				CreatedAt: time.Now().Add(time.Second),
				UpdatedAt: time.Time{},
			},
		},
	}, nil
}

func (im *Im) GetPost(page int, size int) []adapters.PostResult {
	if page < 0 || size < 0 {
		return nil
	}
	offset := page * size
	if offset == 0 {
		offset = size
	}
	var posts []adapters.PostResult
	for _, v := range im.posts {
		posts = append(posts, v)
	}
	sort.Slice(posts, func(i, j int) bool {
		pca, err := time.Parse(posts[i].CreatedAt, time.RFC3339)
		if err != nil {
			fmt.Errorf("cant parse time")
		}
		pua, err := time.Parse(posts[j].CreatedAt, time.RFC3339)
		if err != nil {
			fmt.Errorf("cant parse time")
		}
		return pca.Before(pua)
	})
	return posts
}

func (im *Im) GetPostWithComments(postId int, page int, size int) []adapters.PostWithCommentsResult {
	if page < 0 || size < 0 {
		return nil
	}
	offset := page * size
	if offset == 0 {
		offset = size
	}
	var posts []adapters.PostResult
	var comment []Comment
	for _, v := range im.posts {
		if v.Id == postId {
			posts = append(posts, v)
		}
	}
	for _, v := range im.comment {
		if v.PostId == postId {
			comment = append(comment, v)
		}
	}
	sort.Slice(comment, func(i, j int) bool {
		return comment[i].CreatedAt.Before(comment[j].CreatedAt)
	})

	var res []adapters.PostWithCommentsResult
	for i := range comment {
		var ParId int
		if comment[i].ParId != nil {
			ParId = *comment[i].ParId
		}
		res = append(res, adapters.PostWithCommentsResult{
			PostId:            posts[0].Id,
			PostAuthorId:      posts[0].AuthorId,
			PostCreatedAt:     posts[0].CreatedAt,
			PostUpdatedAt:     posts[0].UpdatedAt,
			PostAllowComments: posts[0].AllowComments,
			CommentId:         comment[i].Id,
			CommentAuthorId:   comment[i].AuthorId,
			CommentParId:      ParId,
			CommentMsg:        comment[i].Msg,
			CommentCreatedAt:  comment[i].CreatedAt.String(),
			CommentUpdatedAt:  comment[i].UpdatedAt.String(),
		})
	}
	return res
}

func (im *Im) CreatePost(userId int, allowComments bool) error {
	return nil
}

func (im *Im) CreateComment(postId int, userId int, parId int, msg string) error {

	return nil
}
