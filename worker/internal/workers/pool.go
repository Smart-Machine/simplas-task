package workers

import (
	"context"
	"log"

	"github.com/Smart-Machine/simplas-project/service/proto"
	"github.com/Smart-Machine/simplas-project/worker/pkg/advertisement"
	"golang.org/x/sync/errgroup"
)

type Pool struct {
	workers         []Worker
	roundRobinIndex int
}

func NewPool(numOfWorkers int, consumerClient proto.ConsumerClient) *Pool {
	workers := []Worker{}
	for i := 0; i < numOfWorkers; i++ {
		workers = append(workers, NewWorker(consumerClient))
	}
	return &Pool{
		workers:         workers,
		roundRobinIndex: 0,
	}
}

func (p *Pool) StartPool(ctx context.Context) *errgroup.Group {
	group, groupCtx := errgroup.WithContext(ctx)
	for i := 0; i < len(p.workers); i++ {
		group.Go(func() error {
			return p.workers[i].StartLoop(groupCtx)
		})
	}
	return group
}

func (p *Pool) SendData(data advertisement.Advertisement) {
	log.Printf("Sending to %d\n", p.roundRobinIndex)
	p.workers[p.roundRobinIndex].SendData(data)
	p.incrRRI()
}

func (p *Pool) incrRRI() {
	if p.roundRobinIndex == len(p.workers)-1 {
		p.roundRobinIndex = 0
	} else {
		p.roundRobinIndex = p.roundRobinIndex + 1
	}
}
