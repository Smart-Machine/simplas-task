package workers

type Worker interface {
	Task()
}

type Producer struct {
	chunks int
}

func (p *Producer) Task() {
	// Question: The first question is how do you section the .json file
	// and send this data via a pool of workers.
}
