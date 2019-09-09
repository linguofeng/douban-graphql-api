package types

import (
	"github.com/graphql-go/graphql"
	"github.com/linguofeng/douban-graphql-api/models"
)

// CommentType test
var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(models.Comment); ok {
					return comment.ID, nil
				}
				return nil, nil
			},
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: UserType,
		},
	},
})
