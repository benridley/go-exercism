package clock

import "fmt"

// Clock is a simple struct to represent time. This is represented by the number of minutes since
type Clock struct {
	minutes int
}

// New creates a Clock struct. Excess minutes are converted to hours.
func New(h int, m int) Clock {
	clock := Clock{
		minutes: ((((h*60 + m) % 1440) + 1440) % 1440),
	}
	return clock
}

// ToString returns a stringified clock
func (clock Clock) ToString() string {
	return fmt.Sprintf("Hour: %d, Minute: %d", clock.minutes/60, clock.minutes%60)
}

// Add Minutes to a clock.
func (clock Clock) Add(min int) Clock {
	clock.minutes = (((clock.minutes + min) % 1440) + 1440) % 1440
	return clock
}

// Subtract minutes from a clock
func (clock Clock) Subtract(min int) Clock {
	return clock.Add(-min)
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.minutes/60, clock.minutes%60)
}
