package db

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/graiendor/day03/internal"
	"log"
	"strconv"
)

var (
	DocId int64
)

func CreateEntry(indexName string, place internal.Place) {
	indexer := CreateIndexer()
	if DocId == 0 {
		DocId = int64(GetLastId(indexName))
	}
	DocId++
	place.ID = DocId
	data, err := json.Marshal(place)

	if err != nil {
		log.Fatalf("Error marshalling data: %s", err)
	}
	log.Printf("Indexing place %d", DocId)
	err = indexer.Add(context.Background(), esutil.BulkIndexerItem{
		Action:     "index",
		Body:       bytes.NewReader(data),
		DocumentID: strconv.Itoa(int(DocId)),
		OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
			log.Println("success")
			log.Printf("%s %s", res.Result, res.DocumentID)
		},
		OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
			log.Println("failure")
			log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
		},
	},
	)
	if err != nil {
		log.Fatalf("Error adding item to indexer: %s", err)
	}
	indexer.Close(context.Background())
}

func GetLastId(s string) int {
	res, err := client.Search(
		client.Search.WithIndex(s),
		client.Search.WithSize(1),
		client.Search.WithSort(
			"id:desc",
		),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	var mapInterface map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&mapInterface); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	if mapInterface["hits"] == nil {
		return 0
	}

	if len(mapInterface["hits"].(map[string]interface{})["hits"].([]interface{})) == 0 {
		return 0
	} else {
		return int(mapInterface["hits"].(map[string]interface{})["hits"].(interface{}).([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})["id"].(float64))
	}
}
