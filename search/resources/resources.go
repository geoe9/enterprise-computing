package resources

import (
	"encoding/json"
	"enterprise-computing/search/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Track struct {
	Id    string
	Audio string
}

type Response struct {
	Id string
}

func SearchTrack(w http.ResponseWriter, r *http.Request) {
	var t Track
	if err := json.NewDecoder(r.Body).Decode(&t); err == nil {
		if apiRes, err := service.IdentifyTrack(t.Audio); err == nil {
			if apiRes.Status == "success" {
				if apiRes.Result.Title != "" {
					w.WriteHeader(200) /* OK */
					json.NewEncoder(w).Encode(Response{Id: apiRes.Result.Title})
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
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/search", SearchTrack).Methods("POST")
	return r
}
