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
	},
})
