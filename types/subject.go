package types

import (
	"github.com/graphql-go/graphql"
	reviewRepo "github.com/linguofeng/douban-graphql-api/douban/review/repository"
	reviewUsecase "github.com/linguofeng/douban-graphql-api/douban/review/usecase"
	subjectRepo "github.com/linguofeng/douban-graphql-api/douban/subject/repository"
	subjectUsecase "github.com/linguofeng/douban-graphql-api/douban/subject/usecase"
	"github.com/linguofeng/douban-graphql-api/models"
)

// AllSubjectType 所有主题类型
var AllSubjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AllSubject",
	Fields: graphql.Fields{
		"showing": &graphql.Field{
			Type: graphql.NewList(SubjectType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c := subjectUsecase.NewSubjectUsecase(subjectRepo.NewHttpSubjectRepository())
				return c.FetchMovieShowing(1, 10)
			},
		},
		"hotGaia": &graphql.Field{
			Type: graphql.NewList(SubjectType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c := subjectUsecase.NewSubjectUsecase(subjectRepo.NewHttpSubjectRepository())
				return c.FetchMovieHotGaia(1, 10)
			},
		},
		"tvHot": &graphql.Field{
			Type: graphql.NewList(SubjectType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c := subjectUsecase.NewSubjectUsecase(subjectRepo.NewHttpSubjectRepository())
				return c.FetchTvHot(1, 10)
			},
		},
		"tvVarietyShow": &graphql.Field{
			Type: graphql.NewList(SubjectType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c := subjectUsecase.NewSubjectUsecase(subjectRepo.NewHttpSubjectRepository())
				return c.FetchTvVarietyShow(1, 10)
			},
		},
		"bookBestseller": &graphql.Field{
			Type: graphql.NewList(SubjectType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c := subjectUsecase.NewSubjectUsecase(subjectRepo.NewHttpSubjectRepository())
				return c.FetchBookBestseller(1, 10)
			},
		},
		"musicSingle": &graphql.Field{
			Type: graphql.NewList(SubjectType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c := subjectUsecase.NewSubjectUsecase(subjectRepo.NewHttpSubjectRepository())
				return c.FetchMusicSingle(1, 10)
			},
		},
	},
})

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
		"originalTitle": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"intro": &graphql.Field{
			Type: graphql.String,
		},
		"year": &graphql.Field{
			Type: graphql.String,
		},
		"image": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				subject := p.Source.(*models.SubjectDetail)
				return subject.Image.Normal, nil
			},
		},
		"genres": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"countries": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"durations": &graphql.Field{
			Type: graphql.NewList(graphql.String),
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
				subject := p.Source.(*models.SubjectDetail)
				r := reviewUsecase.NewReviewUsecase(reviewRepo.NewHttpReviewRepository())
				return r.Fetch(subject.Type, subject.ID)
			},
		},
	},
})
