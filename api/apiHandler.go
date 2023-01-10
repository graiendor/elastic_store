package api

import (
	"encoding/json"
	"fmt"
	"github.com/graiendor/day03/db"
	"github.com/graiendor/day03/internal"
	"log"
	"net/http"
	"strconv"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("page") != "" {
		limit := 10
		w.Header().Set("Content-Type", "application/json")
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 0 {
			writeJsonError(w, r.URL.Query().Get("page"))
		} else {
			serveJson(w, r, limit, page)
		}
	} else {

		w.WriteHeader(http.StatusBadRequest)
	}
}

func serveJson(w http.ResponseWriter, r *http.Request, limit int, page int) {
	places, total, err := db.GetPlaces(limit, page*limit)
	if err != nil {
		log.Fatalf("Error getting places: %s", err)
	}
	if len(places) > 0 {
		data := struct {
			Name   string           `json:"name"`
			Total  int              `json:"total"`
			Places []internal.Place `json:"places"`
		}{
			Name:   "places",
			Total:  total,
			Places: places,
		}
		w.WriteHeader(200)
		WriteJson(w, data)
	} else {
		writeJsonError(w, r.URL.Query().Get("page"))
	}
}

func writeJsonError(w http.ResponseWriter, page string) {
	data := struct {
		Error string `json:"error"`
	}{
		Error: fmt.Sprintf("Invalid 'page' value: %s", page),
	}
	w.WriteHeader(400)
	WriteJson(w, data)
}

func WriteJson(w http.ResponseWriter, data interface{}) {
	var bytes []byte
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshaling json: %s", err)
	}
	w.Write(bytes)
}
