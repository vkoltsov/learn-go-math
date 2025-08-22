package myMath

import (
	"time"
)

func Add(a, b int) int {
	return a + b
}

type Stopwatch struct {
	started time.Time

	Results []time.Duration
	Stopped bool
}

func (s *Stopwatch) Start() {
	if s.Stopped {
		s.Results = make([]time.Duration, 0)
		s.started = time.Now()
	}
	s.Stopped = true
}

func (s *Stopwatch) SaveSplit() {
	if !s.Stopped {
		s.Results = append(s.Results, time.Since(s.started))
	}
}

func (s Stopwatch) GetResults() []time.Duration {
	return s.Results
}
