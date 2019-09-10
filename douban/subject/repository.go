package subject

import (
	"github.com/linguofeng/douban-graphql-api/models"
)

type Repository interface {
	Fetch(stype SubjectType, start int, count int) ([]*models.Subject, error)
	GetById(stype string, id string) (*models.SubjectDetail, error)
}
