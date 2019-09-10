package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/linguofeng/douban-graphql-api/douban/subject"
	"github.com/linguofeng/douban-graphql-api/models"
)

type httpSubjectRepository struct {
	url    string
	apikey string
}

func NewHttpSubjectRepository() subject.Repository {
	return &httpSubjectRepository{
		url:    "https://frodo.douban.com/api/v2",
		apikey: "054022eaeae0b00e0fc068c0c0a2102a",
	}
}

func (h *httpSubjectRepository) Fetch(stype subject.SubjectType, start int, count int) (res []*models.Subject, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/subject_collection/%s/items?start=%d&count=%d&apiKey=%s", h.url, stype, start, count, h.apikey))
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

func (h *httpSubjectRepository) GetById(stype string, id string) (*models.SubjectDetail, error) {
	fmt.Println(fmt.Sprintf("%s/%s/%s?apiKey=%s", h.url, stype, id, h.apikey))
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s?apiKey=%s", h.url, stype, id, h.apikey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	subject := new(models.SubjectDetail)

	err = json.Unmarshal(body, &subject)
	if err != nil {
		return nil, err
	}

	return subject, nil
}
