package db

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/graiendor/day03/internal"
	"log"
)

var (
	client *elasticsearch.Client
)

func Initialize() {
	CreateClient()
	CreateIndex("places")
	CreateEntry("places", internal.Place{
		Name:    "Budapest",
		Address: "Budapest, Hungary",
		Phone:   "+36 1 234 5678",
		Location: []float64{
			47.497912,
			19.040235,
		},
	})

	CreateEntry("places", internal.Place{
		Name:    "London",
		Address: "London, UK",
		Phone:   "+44 20 7946 0532",
		Location: []float64{
			51.507351,
			-0.127758,
		},
	})

	CreateEntry("places", internal.Place{
		Name:    "New York",
		Address: "New York, USA",
		Phone:   "+1 212 938 3100",
		Location: []float64{
			40.712728,
			-74.006015,
		},
	})

	//for i := 0; i < 6004; i++ {
	//	CreateEntry("places", internal.Place{
	//		Name:    "Xaerw",
	//		Address: "Test",
	//		Phone:   "Test",
	//		Location: []float64{
	//			-122.4194,
	//			37.7749,
	//		},
	//	})
	//}
	//
	//for i := 0; i < 6000; i++ {
	//	CreateEntry("places", internal.Place{
	//		Name:    "Andrea",
	//		Address: "Test",
	//		Phone:   "Test",
	//		Location: []float64{
	//			-122.4194,
	//			37.7749,
	//		},
	//	})
	//}
}

func CreateClient() {
	var err error
	client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	log.Println("Client created")
}

func CreateIndexer() esutil.BulkIndexer {
	indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: client,
		Index:  "places",
	})
	return indexer
}
