package entry

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Smart-Machine/simplas-project/worker/pkg/advertisement"
)

type Entry struct {
	Error         error
	Advertisement advertisement.Advertisement
}

type Stream struct {
	stream chan Entry
}

func NewJSONStream() Stream {
	return Stream{
		stream: make(chan Entry),
	}
}

func (s Stream) Watch() <-chan Entry {
	return s.stream
}

func (s Stream) Start(path string) {
	defer close(s.stream)

	file, err := os.Open(path)
	if err != nil {
		s.stream <- Entry{Error: fmt.Errorf("open file: %w", err)}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Error: fmt.Errorf("decode opening delimiter: %w", err)}
		return
	}

	line := 1
	for decoder.More() {
		var advertisement advertisement.Advertisement
		if err := decoder.Decode(&advertisement); err != nil {
			s.stream <- Entry{Error: fmt.Errorf("decode line %d: %w", line, err)}
			return
		}
		s.stream <- Entry{Advertisement: advertisement}
		line++
	}

	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Error: fmt.Errorf("decode closing delimiter: %w", err)}
		return
	}
}
