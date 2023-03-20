package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	KEY = "" /* AUDD.IO API KEY HERE */
	URI = "https://api.audd.io/"
)

type APIResponse struct {
	Status string
	Result struct {
		Title string
	}
}

func IdentifyTrack(audio string) (APIResponse, error) {
	var apiRes APIResponse
	data := url.Values{
		"api_token": {KEY},
		"audio":     {audio},
	}
	if res, err := http.PostForm(URI, data); err == nil {
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(&apiRes); err == nil {
			return apiRes, nil
		}
		return APIResponse{}, errors.New("Service")
	}
	return APIResponse{}, errors.New("Service")
}
