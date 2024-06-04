package gql

import "github.com/graphql-go/graphql"

var Post = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"Id": &graphql.Field{
				Type: graphql.Int,
			},
			"AuthorId": &graphql.Field{
				Type: graphql.Int,
			},
			"CreatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"UpdatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"AllowComments": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var PostWithComments = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PostWithComments",
		Fields: graphql.Fields{
			"PostId": &graphql.Field{
				Type: graphql.Int,
			},
			"PostAuthorId": &graphql.Field{
				Type: graphql.Int,
			},
			"PostCreatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"PostUpdatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"PostAllowComments": &graphql.Field{
				Type: graphql.Boolean,
			},
			"CommentId": &graphql.Field{
				Type: graphql.Int,
			},
			"CommentAuthorId": &graphql.Field{
				Type: graphql.Int,
			},
			"CommentParId": &graphql.Field{
				Type: graphql.Int,
			},
			"CommentMsg": &graphql.Field{
				Type: graphql.String,
			},
			"CommentCreatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"CommentUpdatedAt": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var CreatePost = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CreatePost",
		Fields: graphql.Fields{
			"UserId": &graphql.Field{
				Type: graphql.Int,
			},
			"AllowComments": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var CreateComment = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CreateComment",
		Fields: graphql.Fields{
			"PostId": &graphql.Field{
				Type: graphql.Int,
			},
			"UserId": &graphql.Field{
				Type: graphql.Int,
			},
			"ParId": &graphql.Field{
				Type: graphql.Int,
			},
			"Msg": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
