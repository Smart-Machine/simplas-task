package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Smart-Machine/simplas-project/service/proto"
	"github.com/elastic/go-elasticsearch/v8"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"strings"
)

type ConsumerServer struct {
	elasticClient *elasticsearch.Client
	proto.UnimplementedConsumerServer
}

func (c *ConsumerServer) ConsumeData(_ context.Context, req *proto.ConsumeDataRequest) (*proto.ConsumeDataResponse, error) {
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

type CRUDServer struct {
	elasticClient *elasticsearch.Client
	proto.UnimplementedCRUDServer
}

func (c *CRUDServer) Create(_ context.Context, req *proto.ConsumeDataRequest) (*proto.ConsumeDataResponse, error) {
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

func (c *CRUDServer) GetList(search *wrapperspb.StringValue, stream proto.CRUD_GetListServer) error {
	//for _, advertisement := range
	query := `
		{
			"query": {
				"match": {
					"title": %s
				}
			}
		}
	`

	res, err := c.elasticClient.Search(
		//c.elasticClient.Search.WithBody(strings.NewReader(reqStr)),
		c.elasticClient.Search.WithQuery(fmt.Sprintf(query, search.String())),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	log.Println(res)
	return nil
}

func (c *CRUDServer) GetOne(_ context.Context, id *wrapperspb.Int64Value) (*proto.ConsumeDataResponse, error) {
	res, err := c.elasticClient.Get(
		"advertisements",
		id.String(),
		c.elasticClient.Get.WithRefresh(true),
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

func (c *CRUDServer) Update(_ context.Context, req *proto.UpdateRequest) (*proto.ConsumeDataResponse, error) {
	reqJson, err := json.Marshal(req.Data)
	if err != nil {
		return nil, err
	}
	reqStr := string(reqJson)

	res, err := c.elasticClient.Update(
		"advertisements",
		req.Id.String(),
		strings.NewReader(reqStr),
		c.elasticClient.Update.WithRefresh("true"),
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

func (c *CRUDServer) Delete(_ context.Context, req *wrapperspb.Int64Value) (*wrapperspb.BoolValue, error) {
	res, err := c.elasticClient.Delete(
		"advertisements",
		req.String(),
		c.elasticClient.Delete.WithRefresh("true"),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	log.Println(res.StatusCode)
	return &wrapperspb.BoolValue{Value: true}, nil
}
