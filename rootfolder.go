package main

import (	
	"fmt"
	"encoding/json"
	"github.com/go-resty/resty"
)

type RootFolder struct {
	Id int `json:"id"`
	Path string `json:"path"`
	FreeSpace int `json:"freeSpace"`
}

type RootFolders struct {
	ApiKey string
	Url    string
}

func (r *RootFolders) GetRootFolders() []RootFolder {
	var rootfolders []RootFolder
	requestURL := fmt.Sprintf("%s/api/rootfolder?apikey=%s", r.Url, r.ApiKey)
	resp, err := resty.R().Get(requestURL)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(resp.Body(), &rootfolders)
	if err != nil {
		fmt.Println(err)
	}
	return rootfolders
}
