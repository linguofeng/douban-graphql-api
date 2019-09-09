package repository_test

import (
	"testing"

	"github.com/linguofeng/douban-graphql-api/douban/subject/repository"
)

func TestFetch(t *testing.T) {
	r := repository.NewHttpSubjectRepository()
	r.Fetch(1, 10)
}
