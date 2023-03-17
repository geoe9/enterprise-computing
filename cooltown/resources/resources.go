package resources

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

type Track struct {
	Id    string
	Audio string
}

type TrackAudio struct {
	Audio string
}

func searchTrack(w http.ResponseWriter, r *http.Request) {
	var t Track
	hc := http.Client{}
	if req, err := http.NewRequest("POST", "http://127.0.0.1:3001/search", r.Body); err == nil {
		if res, err := hc.Do(req); err == nil {
			defer res.Body.Close()
			if res.StatusCode != 200 {
				w.WriteHeader(res.StatusCode)
				return
			}
			if err := json.NewDecoder(res.Body).Decode(&t); err == nil {
				id := url.QueryEscape(strings.ReplaceAll(t.Id, " ", "+"))
				if req, err := http.NewRequest("GET", "http://127.0.0.1:3000/tracks/"+id, nil); err == nil {
					if res, err := hc.Do(req); err == nil {
						defer res.Body.Close()
						if res.StatusCode != 200 {
							w.WriteHeader(res.StatusCode)
							return
						}
						if err := json.NewDecoder(res.Body).Decode(&t); err == nil {
							w.WriteHeader(200) /* OK */
							var a TrackAudio
							a.Audio = t.Audio
							json.NewEncoder(w).Encode(a)
						} else {
							w.WriteHeader(500) /* Internal Server Error */
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
		} else {
			w.WriteHeader(500) /* Internal Server Error */
		}
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/cooltown", searchTrack).Methods("POST")
	return r
}
