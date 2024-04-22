package main

import (
	"github.com/Smart-Machine/simplas-project/service"
	"log"
)

func main() {
	err := service.NewServer()
	if err != nil {
		log.Fatal(err)
	}
}
