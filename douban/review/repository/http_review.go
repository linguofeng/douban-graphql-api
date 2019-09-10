package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/linguofeng/douban-graphql-api/douban/review"
	"github.com/linguofeng/douban-graphql-api/models"
)

type httpReviewRepository struct {
	url    string
	apikey string
}

func NewHttpReviewRepository() review.Repository {
	return &httpReviewRepository{
		url:    "https://frodo.douban.com/api/v2",
		apikey: "054022eaeae0b00e0fc068c0c0a2102a",
	}
}

func (h *httpReviewRepository) Fetch(stype string, id string) ([]*models.Review, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s/reviews?apiKey=%s", h.url, stype, id, h.apikey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := new(struct {
		Total   int              `json:"total"`
		Reviews []*models.Review `json:"reviews"`
	})

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data.Reviews, nil
}
