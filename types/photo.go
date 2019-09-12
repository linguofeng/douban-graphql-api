package types

import (
	"github.com/graphql-go/graphql"
)

var PhotoResultType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PhotoResult",
	Fields: graphql.Fields{
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"start": &graphql.Field{
			Type: graphql.Int,
		},
		"count": &graphql.Field{
			Type: graphql.Int,
		},
		"photos": &graphql.Field{
			Type: graphql.NewList(PhotoType),
		},
	},
})

// PhotoType test
var PhotoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Photo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"image": &graphql.Field{
			Type: PhotoImageType,
		},
	},
})

var PhotoImageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PhotoImage",
	Fields: graphql.Fields{
		"small": &graphql.Field{
			Type: PhotoImageDetailType,
		},
		"large": &graphql.Field{
			Type: PhotoImageDetailType,
		},
		"normal": &graphql.Field{
			Type: PhotoImageDetailType,
		},
	},
})

var PhotoImageDetailType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PhotoImageDetail",
	Fields: graphql.Fields{
		"url": &graphql.Field{
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
