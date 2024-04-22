package workers

import (
	"context"
	"log"

	"github.com/Smart-Machine/simplas-project/service/proto"
	"github.com/Smart-Machine/simplas-project/worker/pkg/advertisement"
)

type Worker struct {
	data           chan advertisement.Advertisement
	consumerClient proto.ConsumerClient
}

func NewWorker(consumerClientConn proto.ConsumerClient) Worker {
	return Worker{
		data:           make(chan advertisement.Advertisement),
		consumerClient: consumerClientConn,
	}
}

func (w Worker) StartLoop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			close(w.data)
			return nil
		case d := <-w.data:
			if err := w.processData(ctx, d); err != nil {
				return err
			}
		}
	}
}

func (w Worker) SendData(data advertisement.Advertisement) {
	w.data <- data
}

func (w Worker) processData(ctx context.Context, data advertisement.Advertisement) error {
	res, err := w.consumerClient.ConsumeData(ctx, &proto.ConsumeDataRequest{
		Id:         data.ID,
		Categories: data.Categories,
		Title:      data.Title,
		Type:       data.Type,
		Posted:     data.Posted,
	})
	if err != nil {
		return err
	}
	log.Println(res)
	return nil
}
