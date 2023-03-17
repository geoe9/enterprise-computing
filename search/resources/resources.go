package resources

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

const API_TOKEN = "e279ea64f634ef35e7aa7adf4a97ec42" /* AUDD.IO API HERE PLEASE */

type AudioSnippet struct {
	Audio string
}

type APIResponse struct {
	Status string
	Result struct {
		Title string
	}
}

type SongTitle struct {
	Id string
}

func searchTrack(w http.ResponseWriter, r *http.Request) {
	var input AudioSnippet
	var apiRes APIResponse
	if err := json.NewDecoder(r.Body).Decode(&input); err == nil {
		data := url.Values{
			"api_token": {API_TOKEN},
			"audio":     {input.Audio},
		}
		if res, err := http.PostForm("https://api.audd.io/", data); err == nil {
			defer res.Body.Close()
			if err := json.NewDecoder(res.Body).Decode(&apiRes); err == nil {
				if apiRes.Status == "success" {
					if apiRes.Result.Title != "" {
						w.WriteHeader(200) /* OK */
						json.NewEncoder(w).Encode(SongTitle{apiRes.Result.Title})
					} else {
						w.WriteHeader(404) /* Not Found */
					}
				} else {
					w.WriteHeader(400) /* Bad Request */
				}
			} else {
				w.WriteHeader(500) /* Internal Server Error */
			}
		} else {
			w.WriteHeader(500) /* Internal Server Error */
		}
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/search", searchTrack).Methods("POST")
	return r
}
