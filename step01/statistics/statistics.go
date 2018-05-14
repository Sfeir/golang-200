package statistics

import (
	"time"
)

const (
	statisticsChannelSize = 1000
)

// TODO define a Statistics struct with an uint8 chan, an uint32 counter, a start time and logging period duration
// Statistics is the worker to persist the request statistics
type Statistics struct {
}

// NewStatistics creates a new statistics structure and launches its worker routine
func NewStatistics(loggingPeriod time.Duration) *Statistics {
	// TODO build a new Statistics object with a sized channel, initialized counter and start time
	// TODO and logging period as param

	// TODO launch the run in a separate Go routine in background

	// TODO return the initialized and started object
	return nil
}

// PlusOne is used to send a statistics hit increment
func (sw *Statistics) PlusOne() {
	// TODO push a hit in the statistics channel
}

func (sw *Statistics) run() {
	// TODO build a new time Ticker from the logging period

	// TODO build a infinite loop and the channel selection inside

	// TODO build a first select case from the statistics channel
	// TODO add the hit count to the counter and log it as debug level

	// TODO build a second case on the time Ticker chan
	// TODO retrieve the elapsed time since start
	// TODO log the hit/sec rate
	// TODO reset the counter and the start time
}
