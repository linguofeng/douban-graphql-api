package usecase_test

import (
	"testing"

	"github.com/linguofeng/douban-graphql-api/douban/subject/repository"
	"github.com/linguofeng/douban-graphql-api/douban/subject/usecase"
)

func TestFetch(t *testing.T) {
	r := usecase.NewSubjectUsecase(repository.NewHttpSubjectRepository())
	r.FetchMovieHotGaia(1, 10)
}
