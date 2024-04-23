package service

import (
	"log"
	"net"
	"os"

	"github.com/Smart-Machine/simplas-project/service/proto"
	"github.com/elastic/go-elasticsearch/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func connToElastic() (*elasticsearch.Client, error) {
	config := elasticsearch.Config{
		Addresses: []string{
			"http://64.23.174.193:9200",
		},
		Username: os.Getenv("ELASTIC_USERNAME"), //"elastic",
		Password: os.Getenv("ELASTIC_PASSWORD"), //"Bm4puJHi2C0aR56nhXY5",
	}
	elasticsearchDefaultClient, err := elasticsearch.NewClient(config)
	if err != nil {
		// log.Fatalf("Error creating the client: %s", err)
		return nil, err
	}

	return elasticsearchDefaultClient, nil
}

func seedElastic(elasticClient *elasticsearch.Client) error {
	_, err := elasticClient.Indices.Create("advertisements")
	if err != nil {
		// log.Fatalf("Error creating the index: %s", err)
		return err
	}
	return nil
}

func NewConsumerServer() error {
	elasticClient, err := connToElastic()
	if err != nil {
		return err
	}
	log.Println("Service connected to Elasticsearch")
	log.Println(elasticsearch.Version)
	log.Println(elasticClient.Info())

	err = seedElastic(elasticClient)
	if err != nil {
		return err
	}
	log.Println("Seeded Elasticsearch")

	gRPCListener, err := net.Listen("tcp", ":8000")
	if err != nil {
		return err
	}
	log.Println("Service listening on :8000")

	gRPCServer := grpc.NewServer()
	proto.RegisterConsumerServer(gRPCServer, &ConsumerServer{elasticClient: elasticClient})

	reflection.Register(gRPCServer)

	return gRPCServer.Serve(gRPCListener)
}

func NewCRUDServer() error {
	elasticClient, err := connToElastic()
	if err != nil {
		return err
	}
	log.Println("Service connected to Elasticsearch")
	log.Println(elasticsearch.Version)
	log.Println(elasticClient.Info())

	gRPCListener, err := net.Listen("tcp", ":8001")
	if err != nil {
		return err
	}
	log.Println("API listening on :8001")

	gRPCServer := grpc.NewServer()
	proto.RegisterCRUDServer(gRPCServer, &CRUDServer{elasticClient: elasticClient})

	reflection.Register(gRPCServer)

	return gRPCServer.Serve(gRPCListener)
}
