package service

import (
	"context"
	"encoding/json"
	"github.com/Smart-Machine/simplas-project/service/proto"
	"github.com/elastic/go-elasticsearch/v8"
	"strings"
)

type ConsumerServer struct {
	elasticClient *elasticsearch.Client
	proto.UnimplementedConsumerServer
}

func (c *ConsumerServer) ConsumeData(ctx context.Context, req *proto.ConsumeDataRequest) (*proto.ConsumeDataResponse, error) {
	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	reqStr := string(reqJson)

	res, err := c.elasticClient.Index(
		"advertisements",
		strings.NewReader(reqStr),
		c.elasticClient.Index.WithDocumentID(string(req.Id)),
		c.elasticClient.Index.WithRefresh("true"),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &proto.ConsumeDataResponse{
		StatusCode: int32(res.StatusCode),
		Content:    res.String(),
	}, nil
}
