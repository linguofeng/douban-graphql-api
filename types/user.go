package types

import (
	"github.com/graphql-go/graphql"
	"github.com/linguofeng/douban-graphql-api/models"
)

// UserType test
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "user",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"uid": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.UID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"avatar": &graphql.Field{
			Type: graphql.String,
		},
	},
})
