package subject

import (
	"github.com/linguofeng/douban-graphql-api/models"
)

type Usecase interface {
	Fetch(start int, count int) (res []*models.Subject, err error)
	GetById(stype string, id string) (res *models.SubjectDetail, err error)
}
