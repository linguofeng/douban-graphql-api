package usecase_test

import (
	"fmt"
	"testing"

	"github.com/linguofeng/douban-graphql-api/douban/photo/repository"
	"github.com/linguofeng/douban-graphql-api/douban/photo/usecase"
)

func TestFetch(t *testing.T) {
	r := usecase.NewPhotoUsecase(repository.NewHttpPhotoRepository())
	data, _ := r.Fetch("movie", "26709258")
	fmt.Println(data)
}
