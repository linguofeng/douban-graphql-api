package repository_test

import (
	"fmt"
	"testing"

	"github.com/linguofeng/douban-graphql-api/douban/subject/repository"
)

func TestFetch(t *testing.T) {
	r := repository.NewHttpSubjectRepository()
	r.Fetch("movie_showing", 1, 10)
}

func TestGetById(t *testing.T) {
	r := repository.NewHttpSubjectRepository()
	subject, _ := r.GetById("movie", "26709258")
	fmt.Println(subject)
}
