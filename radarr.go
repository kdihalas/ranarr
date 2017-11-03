package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
)

type Movie struct {
	Title  string  `json:"title"`
	Images []Image `json:"images"`
}

type Radarr struct {
	ApiKey string
	Url    string
}

func (s *Radarr) GetMovies() []Movie {
	var movies []Movie
	requestURL := fmt.Sprintf("%s/api/movie?apikey=%s", s.Url, s.ApiKey)
	resp, err := resty.R().Get(requestURL)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(resp.Body(), &movies)
	if err != nil {
		fmt.Println(err)
	}
	return movies
}
