package service

import (
	"github.com/Smart-Machine/simplas-project/service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewConsumerClient() (proto.ConsumerClient, error) {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return proto.NewConsumerClient(conn), nil
}
