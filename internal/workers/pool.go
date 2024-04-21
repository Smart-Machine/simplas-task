package workers

import "sync"

type Pool struct {
	workers   chan Worker
	waitGroup sync.WaitGroup
}

func NewPool(maxGoroutines int) *Pool {
	pool := Pool{workers: make(chan Worker)}
	pool.waitGroup.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for worker := range pool.workers {
				worker.Task()
			}
			pool.waitGroup.Done()
		}()
	}
	return &pool
}

func (p *Pool) Run(worker Worker) {
	p.workers <- worker
}

func (p *Pool) Shutdown() {
	close(p.workers)
	p.waitGroup.Wait()
}
