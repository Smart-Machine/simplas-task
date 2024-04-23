package main

import (
	"github.com/Smart-Machine/simplas-project/service/pkg/service"
	"log"
)

func main() {
	err := service.NewConsumerServer()
	if err != nil {
		log.Fatal(err)
	}
	err = service.NewCRUDServer()
	if err != nil {
		log.Fatal(err)
	}
}
