package db

import (
	"log"
	"strings"
)

func CreateIndex(indexName string) {
	mapping := `{
		"mappings": {
			"properties": {
				"id": {
					"type": "long"
				},
				"name": {	
					"type": "text"
				},
				"address": {
					"type": "text"
				},
				"phone": {
					"type": "text"	
				},
				"location": {
					"type": "geo_point"
				}
			}
			"settings": {
				"number_of_shards": 1,
				"number_of_replicas": 0
				"max_result_window": 100000
			}
		}
		"aliases": {
			"place": {}
		}
	}`
	found := checkIndex(indexName)
	if !found {
		create, err := client.Indices.Create(indexName, client.Indices.Create.WithBody(strings.NewReader(mapping)))
		if err != nil {
			log.Fatalf("Error creating index: %s", err)
		}
		defer create.Body.Close()
		log.Println("Index created")
	} else {
		log.Println("Index found")
	}
}

func checkIndex(indexName string) bool {
	res, _ := client.Indices.Exists([]string{indexName})
	defer res.Body.Close()
	return res.StatusCode == 200
}
