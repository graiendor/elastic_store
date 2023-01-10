package main

import (
	"fmt"
	"github.com/graiendor/day03/api"
	"github.com/graiendor/day03/html"
	"log"
	"net/http"
)

func run() {
	fmt.Printf("Starting server at port 8888\n")
	fmt.Printf("Visit http://localhost:8888\n")
	fmt.Printf("Press Ctrl+C to stop\n")
	http.HandleFunc("/", html.QueryHandler)
	http.HandleFunc("/api/places", api.ApiHandler)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
