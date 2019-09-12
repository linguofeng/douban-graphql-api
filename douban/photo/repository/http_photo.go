package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/linguofeng/douban-graphql-api/douban/photo"
	"github.com/linguofeng/douban-graphql-api/models"
)

type httpPhotoRepository struct {
	url    string
	apikey string
}

func NewHttpPhotoRepository() photo.Repository {
	return &httpPhotoRepository{
		url:    "https://frodo.douban.com/api/v2",
		apikey: "054022eaeae0b00e0fc068c0c0a2102a",
	}
}

func (h *httpPhotoRepository) Fetch(stype string, id string) (*models.PhotoResp, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s/photos?apiKey=%s", h.url, stype, id, h.apikey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	photoResp := new(models.PhotoResp)

	err = json.Unmarshal(body, &photoResp)
	if err != nil {
		return nil, err
	}

	return photoResp, nil
}
