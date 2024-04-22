package main

import (
	"context"
	"log"
	"time"

	"github.com/Smart-Machine/simplas-project/service"
	"github.com/Smart-Machine/simplas-project/worker/internal/workers"
	"github.com/Smart-Machine/simplas-project/worker/pkg/entry"
)

const (
	dataFilepath  = "/home/smart0machine/Projects/simplas-task/task/data.json"
	numGoroutines = 3
)

func main() {

	client, err := service.NewConsumerClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	pool := workers.NewPool(numGoroutines, client)
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
