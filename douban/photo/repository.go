package photo

import (
	"github.com/linguofeng/douban-graphql-api/models"
)

type Repository interface {
	Fetch(stype string, id string) (*models.PhotoResp, error)
}
