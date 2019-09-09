package types

import (
	"github.com/graphql-go/graphql"
	"github.com/linguofeng/douban-graphql-api/models"
)

// SubjectType test
var SubjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subject",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if subject, ok := p.Source.(models.Subject); ok {
					return subject.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"cover": &graphql.Field{
			Type: CoverType,
		},
		"rating": &graphql.Field{
			Type: RatingType,
		},
	},
})
