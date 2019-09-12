package usecase

import (
	"github.com/linguofeng/douban-graphql-api/douban/photo"
	"github.com/linguofeng/douban-graphql-api/models"
)

type photoUsecase struct {
	repository photo.Repository
}

func NewPhotoUsecase(repository photo.Repository) photo.Usecase {
	return &photoUsecase{
		repository: repository,
	}
}

func (s *photoUsecase) Fetch(stype string, id string) (*models.PhotoResp, error) {
	return s.repository.Fetch(stype, id)
}
