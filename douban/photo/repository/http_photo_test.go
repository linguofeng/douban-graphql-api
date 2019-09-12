package repository_test

import (
	"fmt"
	"testing"

	"github.com/linguofeng/douban-graphql-api/douban/photo/repository"
)

func TestFetch(t *testing.T) {
	r := repository.NewHttpPhotoRepository()
	data, _ := r.Fetch("movie", "26709258")
	fmt.Println(data)
}
