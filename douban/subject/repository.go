package subject

import (
	"github.com/linguofeng/douban-graphql-api/models"
)

type Repository interface {
	Fetch(start int, count int) (res []*models.Subject, err error)
}
