package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
)

type Movie struct {
	Title  string  `json:"title"`
	TitleSlug string `json:"titleSlug"`
	Images []Image `json:"images"`
	Year   int	   `json:"year"`
	TmdbId int `json:"tmdbId"`
	QualityProfileId int `json:"qualityProfileId"`
	RootFolderPath  string `json:"rootFolderPath"`
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

func (s *Radarr) SearchForMovies(term string) []Movie {
	var movies []Movie
	requestURL := fmt.Sprintf("%s/api/movies/lookup?term=%s&apikey=%s", s.Url, term, s.ApiKey)
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

func (s *Radarr) GetMovie(tmdbId int) Movie {
	var movie Movie
	requestURL := fmt.Sprintf("%s/api/movies/lookup/tmdb?tmdbid=%d&apikey=%s", s.Url, tmdbId, s.ApiKey)
	resp, err := resty.R().Get(requestURL)

	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(resp.Body(), &movie)
	if err != nil {
		fmt.Println(err)
	}
	return movie
}

func (s *Radarr) DownloadMovie(tmdbId int) error {
	rootFolder := RootFolders{
		ApiKey: s.ApiKey,
		Url: s.Url,
	}

	firstRootFolder := rootFolder.GetRootFolders()[0]

	movie := s.GetMovie(tmdbId)
	movie.QualityProfileId = 1
	movie.RootFolderPath = firstRootFolder.Path

	requestURL := fmt.Sprintf("%s/api/movie?apikey=%s", s.Url, s.ApiKey)
	resp, err := resty.R().SetBody(movie).Post(requestURL)
	if err != nil {
		return err
	}
	if resp.StatusCode() == 200 {
		return nil
	}
	return nil
}
