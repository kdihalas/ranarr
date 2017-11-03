package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
)

type Serie struct {
	Title  string  `json:"title"`
	TitleSlug string `json:"titleSlug"`
	Images []Image `json:"images"`
	Year   int	   `json:"year"`
	TvdbId int `json:"tvdbId"`
	QualityProfileId int `json:"qualityProfileId"`
	Seasons []Season `json:"seasons"`
	RootFolderPath  string `json:"rootFolderPath"`
}

type Season struct {
	SeasonNumber int `json:"seasonNumber"`
	Monitored bool `json:"monitored"`
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

func (s *Sonarr) SearchForSeries(term string) []Serie {
	var series []Serie
	requestURL := fmt.Sprintf("%s/api/series/lookup?term=%s&apikey=%s", s.Url, term, s.ApiKey)
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

func (s *Sonarr) DownloadSerie(tvdbId int) error {
	rootFolder := RootFolders{
		ApiKey: s.ApiKey,
		Url: s.Url,
	}

	firstRootFolder := rootFolder.GetRootFolders()[0]

	serie := s.SearchForSeries(fmt.Sprintf("tvdb:%d", tvdbId))[0]
	serie.QualityProfileId = 1
	serie.RootFolderPath = firstRootFolder.Path

	requestURL := fmt.Sprintf("%s/api/series?apikey=%s", s.Url, s.ApiKey)
	resp, err := resty.R().SetBody(serie).Post(requestURL)
	if err != nil {
		return err
	}
	if resp.StatusCode() == 200 {
		return nil
	}
	return nil
}
