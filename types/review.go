package types

import (
	"github.com/graphql-go/graphql"
)

// ReivewType test
var ReivewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Reivew",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: UserType,
		},
		"rating": &graphql.Field{
			Type: RatingType,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"commentsCount": &graphql.Field{
			Type: graphql.Int,
		},
		"linkersCount": &graphql.Field{
			Type: graphql.Int,
		},
		"shareCount": &graphql.Field{
			Type: graphql.Int,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"HTML": &graphql.Field{
			Type: graphql.String,
		},
		"subject": &graphql.Field{
			Type: SubjectType,
		},
	},
})
