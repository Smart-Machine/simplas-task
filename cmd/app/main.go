package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

const (
	dataFilepath  = "/home/smart0machine/Projects/simplas-task/task/data.json"
	numGoroutines = 3
)

func main() {
	// stream := entry.NewJSONStream()
	// for i := 0; i < numGoroutines; i++ {
	// 	go func() {
	// 		for data := range stream.Watch() {
	// 			if data.Error != nil {
	// 				log.Println(data.Error)
	// 			}
	// 			log.Printf("%v\n", data.Advertisement)
	// 			time.Sleep(time.Second * 5)
	// 		}
	// 	}()
	// }

	// stream.Start(dataFilepath)

	elasticsearchDefaultClient, _ := elasticsearch.NewDefaultClient()
	log.Println(elasticsearch.Version)
	log.Println(elasticsearchDefaultClient.Info())

	elasticsearchDefaultClient.Indices.Create("ads")
}
