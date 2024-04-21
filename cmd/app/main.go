package main

import (
	"context"
	"log"
	"time"

	"github.com/Smart-Machine/simplas-project/internal/workers"
	"github.com/Smart-Machine/simplas-project/pkg/entry"
)

const (
	dataFilepath  = "/home/smart0machine/Projects/simplas-task/task/data.json"
	numGoroutines = 3
)

func main() {
	// config := elasticsearch.Config{
	// 	Addresses: []string{
	// 		"http://64.23.174.193:9200",
	// 	},
	// 	Username: "elastic",
	// }
	// elasticsearchDefaultClient, err := elasticsearch.NewClient(config)
	// if err != nil {
	// 	log.Fatalf("Error creating the client: %s", err)
	// }

	// log.Println(elasticsearch.Version)
	// log.Println(elasticsearchDefaultClient.Info())

	// _, err = elasticsearchDefaultClient.Indices.Create("kitty")
	// if err != nil {
	// 	log.Fatalf("Error creating the index: %s", err)
	// }

	ctx, cancel := context.WithCancel(context.Background())
	pool := workers.NewPool(numGoroutines)
	errgroup := pool.StartPool(ctx)
	stream := entry.NewJSONStream()
	go func() {
		for data := range stream.Watch() {
			if data.Error != nil {
				log.Println(data.Error)
			}
			pool.SendData(data.Advertisement)
			time.Sleep(time.Second * 5)
		}
	}()

	stream.Start(dataFilepath)
	cancel()
	errgroup.Wait()
}
