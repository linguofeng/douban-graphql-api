package types

import (
	"github.com/graphql-go/graphql"
	subjectRepo "github.com/linguofeng/douban-graphql-api/douban/subject/repository"
	subjectUsecase "github.com/linguofeng/douban-graphql-api/douban/subject/usecase"
)

// QueryType test
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"allSubjects": &graphql.Field{
			Type: AllSubjectType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return new(struct{}), nil
			},
		},
		"subject": &graphql.Field{
			Type: SubjectDetailType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.ID,
					Description: "主题ID",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c := subjectUsecase.NewSubjectUsecase(subjectRepo.NewHttpSubjectRepository())
				return c.GetById("movie", p.Args["id"].(string))
			},
		},
	},
})
