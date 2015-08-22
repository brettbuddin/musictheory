package interval

import (
	"testing"
)

var intervalTests = []struct {
	typeFunc          func(int) Interval
	distance          int
	expectedOctaves   int
	expectedDiatonic  int
	expectedChromatic int
}{
	{Perfect, Unison, 0, 0, 0},
	{Perfect, Second, 0, 1, 2},
	{Perfect, Third, 0, 2, 4},
	{Perfect, Fourth, 0, 3, 5},
	{Perfect, Fifth, 0, 4, 7},
	{Perfect, Sixth, 0, 5, 9},
	{Perfect, Seventh, 0, 6, 11},
	{Perfect, Octave, 1, 0, 0},

	{Major, Unison, 0, 0, 0},
	{Major, Second, 0, 1, 2},
	{Major, Third, 0, 2, 4},
	{Major, Fourth, 0, 3, 5},
	{Major, Fifth, 0, 4, 7},
	{Major, Sixth, 0, 5, 9},
	{Major, Seventh, 0, 6, 11},
	{Major, Octave, 1, 0, 0},

	{Augmented, Second, 0, 1, 3},
	{Augmented, Third, 0, 2, 5},
	{Augmented, Fourth, 0, 3, 6},
	{Augmented, Fifth, 0, 4, 8},
	{Augmented, Sixth, 0, 5, 10},
	{Augmented, Seventh, 0, 6, 12}, // TODO: Look into this chroma
	{Augmented, Octave, 1, 0, 1},

	{Diminished, Second, 0, 1, 0},
	{Diminished, Third, 0, 2, 2},
	{Diminished, Fourth, 0, 3, 4},
	{Diminished, Fifth, 0, 4, 6},
	{Diminished, Sixth, 0, 5, 7},
	{Diminished, Seventh, 0, 6, 9},
	{Diminished, Octave, 1, 0, -1},

	{Minor, Third, 0, 2, 3},
	{Minor, Fifth, 0, 4, 6},
	{Minor, Seventh, 0, 6, 10},
}

func TestIntervals(test *testing.T) {
	for i, t := range intervalTests {
		actual := t.typeFunc(t.distance)

		if actual.octaves != t.expectedOctaves ||
			actual.diatonic != t.expectedDiatonic ||
			actual.chromatic != t.expectedChromatic {

			test.Errorf("index=%d actual=%d expected=(octaves=%d diatonic=%d chromatic=%d)",
				i,
				actual,
				t.expectedOctaves,
				t.expectedDiatonic,
				t.expectedChromatic)
		}
	}
}
