package gql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"ozon/internal/adapters"
)

type Resolver struct {
	db adapters.Storage
}

func (r *Resolver) PostResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	if !ok {
		return nil, fmt.Errorf("page dont found")
	}
	size, ok := p.Args["size"].(int)
	if ok {
		posts := r.db.GetPost(page, size)
		return posts, nil
	}

	return nil, nil
}

func (r *Resolver) PostWithCommentsResolver(p graphql.ResolveParams) (interface{}, error) {
	postId, ok := p.Args["postid"].(int)
	if !ok {
		return nil, fmt.Errorf("postId dont found")
	}
	page, ok := p.Args["page"].(int)
	if !ok {
		return nil, fmt.Errorf("page dont found")
	}
	size, ok := p.Args["size"].(int)
	if ok {
		posts := r.db.GetPostWithComments(postId, page, size)
		return posts, nil
	}

	return nil, nil
}

func (r *Resolver) CreatePostResolver(p graphql.ResolveParams) (interface{}, error) {
	userId, ok := p.Args["userid"].(int)
	if !ok {
		return nil, fmt.Errorf("userid dont found")
	}
	allowComments, ok := p.Args["allowcomments"].(bool)
	if ok {
		createPost := r.db.CreatePost(userId, allowComments)
		return createPost, nil
	}

	return nil, nil
}

func (r *Resolver) CreateCommentResolver(p graphql.ResolveParams) (interface{}, error) {
	postId, ok := p.Args["postid"].(int)
	if !ok {
		return nil, fmt.Errorf("postId dont found")
	}
	userId, ok := p.Args["userid"].(int)
	if !ok {
		return nil, fmt.Errorf("userid dont found")
	}
	parId, ok := p.Args["parid"].(int)
	if !ok {
		return nil, fmt.Errorf("msg dont found")
	}
	msg, ok := p.Args["msg"].(string)
	if ok {
		createComment := r.db.CreateComment(postId, userId, parId, msg)
		return createComment, nil
	}

	return nil, nil
}
