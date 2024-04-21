package workers

import (
	"context"
	"log"

	"github.com/Smart-Machine/simplas-project/pkg/advertisement"
)

type Worker struct {
	data chan advertisement.Advertisement
}

func (w Worker) StartLoop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			log.Println("loop closed correctly")
			close(w.data)
			return nil
		case d := <-w.data:
			log.Println("We are here")
			if err := w.processData(d); err != nil {
				return err
			}
		}
	}
}

func (w Worker) SendData(data advertisement.Advertisement) {
	w.data <- data
}

func (w Worker) processData(data advertisement.Advertisement) error {
	// sending data to elastic
	log.Println(data)
	return nil
}
