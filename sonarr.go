package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
)

type Serie struct {
	Title  string  `json:"title"`
	Images []Image `json:"images"`
}

type Image struct {
	CoverType string `json:"coverType"`
	URL       string `json:"url"`
}

type Sonarr struct {
	ApiKey string
	Url    string
}

func (s *Sonarr) GetSeries() []Serie {
	var series []Serie
	requestURL := fmt.Sprintf("%s/api/series?apikey=%s", s.Url, s.ApiKey)
	resp, err := resty.R().Get(requestURL)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(resp.Body(), &series)
	if err != nil {
		fmt.Println(err)
	}
	return series
}
