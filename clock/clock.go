package clock

import (
	"time"
)

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func NewClock() Clock {
	return &realClock{}
}

func (clock *realClock) Now() time.Time {
	return time.Now()
}

type mockClock struct {
	currentTime time.Time
}

func NewMockClock(time time.Time) Clock {
	return &mockClock{
		currentTime: time,
	}
}

func (mockClock *mockClock) Now() time.Time {
	return mockClock.currentTime
}
