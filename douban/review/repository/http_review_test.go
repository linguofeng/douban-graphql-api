package repository_test

import (
	"fmt"
	"testing"

	"github.com/linguofeng/douban-graphql-api/douban/review/repository"
)

func TestFetch(t *testing.T) {
	r := repository.NewHttpReviewRepository()
	reviews, _ := r.Fetch("movie", "26709258")
	fmt.Println(reviews)
}
