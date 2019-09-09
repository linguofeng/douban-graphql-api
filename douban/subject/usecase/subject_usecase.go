package usecase

import (
	"github.com/linguofeng/douban-graphql-api/douban/subject"
	"github.com/linguofeng/douban-graphql-api/models"
)

type subjectUsecase struct {
	repository subject.Repository
}

func NewSubjectUsecase(repository subject.Repository) subject.Usecase {
	return &subjectUsecase{
		repository: repository,
	}
}

func (s *subjectUsecase) Fetch(start int, count int) (res []*models.Subject, err error) {
	return s.repository.Fetch(start, count)
}
