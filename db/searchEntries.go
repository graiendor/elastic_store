package db

import (
	"context"
	"encoding/json"
	"github.com/graiendor/day03/internal"
	"log"
)

func GetPlaces(limit int, offset int) ([]internal.Place, int, error) {
	// Search for the indexed documents
	//
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("places"),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
		client.Search.WithSize(limit),
		client.Search.WithFrom(offset),
		client.Search.WithSort(
			"id:asc",
		),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	var mapInterface map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&mapInterface); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	var places []internal.Place
	for _, hit := range mapInterface["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var place internal.Place
		source := hit.(map[string]interface{})["_source"]
		sourceBytes, _ := json.Marshal(source)
		err := json.Unmarshal(sourceBytes, &place)
		if err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}
		places = append(places, place)
	}

	return places, int(mapInterface["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)), nil
}
