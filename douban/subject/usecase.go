package subject

import (
	"github.com/linguofeng/douban-graphql-api/models"
)

type Usecase interface {
	FetchMovieShowing(start int, count int) (res []*models.Subject, err error)
	FetchMovieHotGaia(start int, count int) (res []*models.Subject, err error)
	FetchTvHot(start int, count int) (res []*models.Subject, err error)
	FetchTvVarietyShow(start int, count int) (res []*models.Subject, err error)
	FetchBookBestseller(start int, count int) (res []*models.Subject, err error)
	FetchMusicSingle(start int, count int) (res []*models.Subject, err error)
	GetById(stype string, id string) (res *models.SubjectDetail, err error)
}
