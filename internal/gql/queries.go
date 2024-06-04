package gql

import (
	"github.com/graphql-go/graphql"
	"ozon/internal/adapters"
)

type Root struct {
	Query *graphql.Object
}

func NewRoot(db adapters.Storage) *Root {

	resolver := Resolver{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"posts": &graphql.Field{

						Type: graphql.NewList(Post),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.PostResolver,
					},
					"postwithcomment": &graphql.Field{
						Type: graphql.NewList(PostWithComments),
						Args: graphql.FieldConfigArgument{
							"postid": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.PostWithCommentsResolver,
					},
					"createpost": &graphql.Field{
						Type: CreatePost,
						Args: graphql.FieldConfigArgument{
							"userid": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"allowcomments": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Boolean),
							},
						},
						Resolve: resolver.CreatePostResolver,
					},
					"createcomment": &graphql.Field{
						Type: CreateComment,
						Args: graphql.FieldConfigArgument{
							"postid": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"userid": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"parid": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"msg": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.CreateCommentResolver},
				},
			},
		),
	}
	return &root
}
