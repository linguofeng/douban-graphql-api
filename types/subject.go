package types

import (
	"github.com/graphql-go/graphql"
	"github.com/linguofeng/douban-graphql-api/models"
	reviewRepo "github.com/linguofeng/douban-graphql-api/douban/review/repository"
	reviewUsecase "github.com/linguofeng/douban-graphql-api/douban/review/usecase"
)

// SubjectType test
var SubjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subject",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if subject, ok := p.Source.(*models.Subject); ok {
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

// SubjectDetailType test
var SubjectDetailType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SubjectDetail",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if subject, ok := p.Source.(*models.SubjectDetail); ok {
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
		"reviews": &graphql.Field{
			Type: graphql.NewList(ReivewType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				r := reviewUsecase.NewReviewUsecase(reviewRepo.NewHttpReviewRepository())
				return r.Fetch("movie", "26709258")
			},
		},
	},
})
