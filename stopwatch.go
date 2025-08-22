package myMath

import (
	"time"
)

// Stopwatch is a utility to measure elapsed time for operations.
type Stopwatch struct {
	started time.Time

	Results []time.Duration
}

// Start initializes the stopwatch and starts timing.
func (s *Stopwatch) Start() {
	s.Results = make([]time.Duration, 0)
	s.started = time.Now()
}

// SaveSplit records the elapsed time since the stopwatch was started
func (s *Stopwatch) SaveSplit() {
	s.Results = append(s.Results, time.Since(s.started))
}

// GetResults returns the recorded elapsed times.
// It returns a slice of time.Duration representing the elapsed times for each split.
func (s Stopwatch) GetResults() []time.Duration {
	return s.Results
}
