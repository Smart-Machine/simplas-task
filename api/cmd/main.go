package main

import (
	"github.com/Smart-Machine/simplas-project/service/pkg/service"
	"log"
)

func main() {
	client, err := service.NewCRUDClient()
	if err != nil {
		log.Fatal(err)
	}

}
