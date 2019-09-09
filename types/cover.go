package types

import (
	"github.com/graphql-go/graphql"
)

// CoverType test
var CoverType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Cover",
	Fields: graphql.Fields{
		"url": &graphql.Field{
			Type: graphql.String,
		},
		"shape": &graphql.Field{
			Type: graphql.String,
		},
		"width": &graphql.Field{
			Type: graphql.Int,
		},
		"height": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
