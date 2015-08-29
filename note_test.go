package mt

import "testing"

func TestNoteDuration(test *testing.T) {
	data := []struct {
		duration Duration
		unit     Duration
		bpm      int
		expected float64
	}{
		{Whole, Quarter, 60, 4},
		{Quarter, Quarter, 60, 1},
		{Eighth, Eighth, 60, 0.25},
		{Whole, Eighth, 60, 2},
	}

	for i, t := range data {
		actual := NewNote(C, 0, Natural, t.duration).Seconds(t.unit, t.bpm)

		if actual != t.expected {
			test.Errorf("index=%d actual=%d expected=%s", i, actual, t.expected)
		}
	}
}
