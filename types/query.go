package types

import (
	"github.com/graphql-go/graphql"
)

// QueryType test
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"subjects": &graphql.Field{
			Type: graphql.NewList(SubjectType),
		},
	},
})
