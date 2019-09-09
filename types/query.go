package types

import (
	"github.com/graphql-go/graphql"
	"github.com/linguofeng/douban-graphql-api/douban/subject/repository"
	"github.com/linguofeng/douban-graphql-api/douban/subject/usecase"
)

// QueryType test
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"subjects": &graphql.Field{
			Type: graphql.NewList(SubjectType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c := usecase.NewSubjectUsecase(repository.NewHttpSubjectRepository())
				return c.Fetch(1, 10)
			},
		},
	},
})
