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

func (s *subjectUsecase) FetchMovieShowing(start int, count int) ([]*models.Subject, error) {
	return s.repository.Fetch(subject.MovieShowing, start, count)
}

func (s *subjectUsecase) FetchMovieHotGaia(start int, count int) ([]*models.Subject, error) {
	return s.repository.Fetch(subject.MovieHotGaia, start, count)
}

func (s *subjectUsecase) FetchTvHot(start int, count int) ([]*models.Subject, error) {
	return s.repository.Fetch(subject.TvHot, start, count)
}

func (s *subjectUsecase) FetchTvVarietyShow(start int, count int) ([]*models.Subject, error) {
	return s.repository.Fetch(subject.TvVarietyShow, start, count)
}

func (s *subjectUsecase) FetchBookBestseller(start int, count int) ([]*models.Subject, error) {
	return s.repository.Fetch(subject.BookBestseller, start, count)
}

func (s *subjectUsecase) FetchMusicSingle(start int, count int) ([]*models.Subject, error) {
	return s.repository.Fetch(subject.MusicSingle, start, count)
}

func (s *subjectUsecase) GetById(stype string, id string) (*models.SubjectDetail, error) {
	return s.repository.GetById(stype, id)
}
