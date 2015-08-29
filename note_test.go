package mt

import (
	"testing"
	"time"
)

func TestNoteDuration(test *testing.T) {
	data := []struct {
		duration Duration
		unit     Duration
		bpm      int
		expected time.Duration
	}{
		{D1, D4, 60, time.Duration(4) * time.Second},
		{D4, D4, 60, time.Duration(1) * time.Second},
		{D8, D8, 60, time.Duration(250) * time.Millisecond},
		{D1, D8, 60, time.Duration(2) * time.Second},
	}

	for i, t := range data {
		actual := NewNote(C, 0, Natural, t.duration).Time(t.unit, t.bpm)

		if actual != t.expected {
			test.Errorf("index=%d actual=%d expected=%s", i, actual, t.expected)
		}
	}
}
