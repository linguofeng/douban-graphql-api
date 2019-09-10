package subject

import (
	"github.com/linguofeng/douban-graphql-api/models"
)

type Repository interface {
	Fetch(start int, count int) ([]*models.Subject, error)
	GetById(stype string, id string) (*models.SubjectDetail, error)
}
