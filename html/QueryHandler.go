package html

import (
	"fmt"
	"github.com/graiendor/day03/db"
	"github.com/graiendor/day03/internal"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type Store interface {
	// GetPlaces returns a list of items, a total number of hits and (or) an error in case of one
	GetPlaces(limit int, offset int) ([]internal.Place, int, error)
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("page") != "" {
		limit := 10
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 0 {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Invalid 'page' value: %s", r.URL.Query().Get("page"))
		} else {
			serveTemplate(w, r, limit, page)
		}
	} else {
		fmt.Fprintf(w, "You searched for nothing")
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request, limit int, page int) {
	lp := filepath.Join("html/static", "index.html")
	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}
	templ := template.Must(tmpl, err)
	places, total, err := db.GetPlaces(limit, page*limit)
	if err != nil {
		log.Fatalf("Error getting places: %s", err)
	}
	if page <= total/limit {
		data := struct {
			Places []internal.Place
			Total  int
			Prev   int
			Next   int
			Last   int
		}{
			Places: places,
			Total:  total,
			Prev:   page - 1,
			Next:   page + 1,
			Last:   total / limit,
		}
		err = templ.Execute(w, data)
		if err != nil {
			log.Fatalf("Error executing template: %s", err)
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Invalid 'page' value: %s", r.URL.Query().Get("page"))
	}
}
