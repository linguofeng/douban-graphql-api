package usecase_test

import (
	"fmt"
	"testing"

	"github.com/linguofeng/douban-graphql-api/douban/review/repository"
	"github.com/linguofeng/douban-graphql-api/douban/review/usecase"
)

func TestFetch(t *testing.T) {
	r := usecase.NewReviewUsecase(repository.NewHttpReviewRepository())
	reviews, _ := r.Fetch("movie", "26709258")
	fmt.Println(reviews)
}
