package types

import (
	"github.com/graphql-go/graphql"
)

// RatingType test
var RatingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Rating",
	Fields: graphql.Fields{
		"count": &graphql.Field{
			Type: graphql.Int,
		},
		"max": &graphql.Field{
			Type: graphql.Int,
		},
		"value": &graphql.Field{
			Type: graphql.Float,
		},
	},
})
