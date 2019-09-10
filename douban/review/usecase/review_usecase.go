package usecase

import (
	"github.com/linguofeng/douban-graphql-api/douban/review"
	"github.com/linguofeng/douban-graphql-api/models"
)

type reviewUsecase struct {
	repository review.Repository
}

func NewReviewUsecase(repository review.Repository) review.Usecase {
	return &reviewUsecase{
		repository: repository,
	}
}

func (s *reviewUsecase) Fetch(stype string, id string) ([]*models.Review, error) {
	return s.repository.Fetch(stype, id)
}
