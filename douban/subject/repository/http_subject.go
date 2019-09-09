package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/linguofeng/douban-graphql-api/douban/subject"
	"github.com/linguofeng/douban-graphql-api/models"
)

type httpSubjectRepository struct {
	url string
}

func NewHttpSubjectRepository() subject.Repository {
	return &httpSubjectRepository{
		url: "https://frodo.douban.com/api/v2/subject_collection/movie_showing/items?start=0&count=20&apiKey=054022eaeae0b00e0fc068c0c0a2102a",
	}
}

func (h *httpSubjectRepository) Fetch(start int, count int) (res []*models.Subject, err error) {
	resp, err := http.Get(h.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := new(struct {
		Count    int               `json:"count"`
		Start    int               `json:"start"`
		Total    int               `json:"total"`
		Subjects []*models.Subject `json:"subject_collection_items"`
	})

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data.Subjects, nil
}
