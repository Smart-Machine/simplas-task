package main

import (
	"context"
	"github.com/Smart-Machine/simplas-project/service/pkg/service"
	"github.com/Smart-Machine/simplas-project/worker/internal/workers"
	"github.com/Smart-Machine/simplas-project/worker/pkg/entry"
	"log"
)

const (
	dataFilepath = "./data.json"
	numOfWorkers = 10
)

func main() {

	client, err := service.NewConsumerClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	pool := workers.NewPool(numOfWorkers, client)
	errgroup := pool.StartPool(ctx)
	stream := entry.NewJSONStream()
	go func() {
		for data := range stream.Watch() {
			if data.Error != nil {
				log.Println(data.Error)
			}
			pool.SendData(data.Advertisement)
			//time.Sleep(time.Second * 1)
		}
	}()

	stream.Start(dataFilepath)
	cancel()
	errgroup.Wait()
}
