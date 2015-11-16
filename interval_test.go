package mt

import (
	"testing"
)

func TestIntervals(test *testing.T) {
	data := []struct {
		typeFunc          IntervalFunc
		step              int
		expectedOctaves   int
		expectedDiatonic  int
		expectedChromatic int
	}{
		{Perfect, 1, 0, 0, 0},
		{Perfect, 2, 0, 1, 2},
		{Perfect, 3, 0, 2, 4},
		{Perfect, 4, 0, 3, 5},
		{Perfect, 5, 0, 4, 7},
		{Perfect, 6, 0, 5, 9},
		{Perfect, 7, 0, 6, 11},
		{Perfect, 8, 1, 0, 0},

		{Major, 1, 0, 0, 0},
		{Major, 2, 0, 1, 2},
		{Major, 3, 0, 2, 4},
		{Major, 4, 0, 3, 5},
		{Major, 5, 0, 4, 7},
		{Major, 6, 0, 5, 9},
		{Major, 7, 0, 6, 11},
		{Major, 8, 1, 0, 0},

		{Augmented, 2, 0, 1, 3},
		{Augmented, 3, 0, 2, 5},
		{Augmented, 4, 0, 3, 6},
		{Augmented, 5, 0, 4, 8},
		{Augmented, 6, 0, 5, 10},
		{Augmented, 7, 0, 6, 12},
		{Augmented, 8, 1, 0, 1},

		{Diminished, 2, 0, 1, 0},
		{Diminished, 3, 0, 2, 2},
		{Diminished, 4, 0, 3, 4},
		{Diminished, 5, 0, 4, 6},
		{Diminished, 6, 0, 5, 7},
		{Diminished, 7, 0, 6, 9},
		{Diminished, 8, 1, 0, -1},

		{Minor, 3, 0, 2, 3},
		{Minor, 5, 0, 4, 6},
		{Minor, 7, 0, 6, 10},
	}

	for i, t := range data {
		actual := t.typeFunc(t.step)

		if actual.Octaves != t.expectedOctaves ||
			actual.Diatonic != t.expectedDiatonic ||
			actual.Chromatic != t.expectedChromatic {

			test.Errorf("index=%d actual=%d expected=(octaves=%d diatonic=%d chromatic=%d)",
				i,
				actual,
				t.expectedOctaves,
				t.expectedDiatonic,
				t.expectedChromatic)
		}
	}
}

func TestIntervalQuality(test *testing.T) {
	data := []struct {
		input    Interval
		expected Quality
	}{
		{Perfect(5), Quality{PerfectType, 0}},
		{Major(2), Quality{MajorType, 0}},
		{Minor(3), Quality{MinorType, 0}},
		{Major(-12), Quality{MinorType, 0}},
		{Augmented(1), Quality{AugmentedType, 1}},
		{DoublyAugmented(1), Quality{AugmentedType, 2}},
		{DoublyDiminished(1), Quality{DiminishedType, 2}},
	}

	for i, t := range data {
		actual := t.input.Quality()
		if !actual.Eq(t.expected) {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

func TestTranspose(test *testing.T) {
	data := []struct {
		initial, interval, expected Interval
	}{
		{Interval{0, 0, 0}, Major(2), Interval{0, 1, 2}},
		{Interval{0, 0, 0}, Major(3), Interval{0, 2, 4}},
		{Interval{0, 0, 0}, Minor(3), Interval{0, 2, 3}},
		{Interval{0, 0, 0}, Augmented(1), Interval{0, 0, 1}},
		{Interval{0, 1, 2}, Augmented(4), Interval{0, 4, 8}},
		{Interval{0, 6, 11}, Minor(3), Interval{1, 1, 2}},
		{Interval{0, 6, 11}, Diminished(5).Negate(), Interval{0, 2, 5}},
	}

	for i, t := range data {
		actual := t.initial.Transpose(t.interval).(Interval)
		if actual.Octaves != t.expected.Octaves ||
			actual.Diatonic != t.expected.Diatonic ||
			actual.Chromatic != t.expected.Chromatic {

			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

func TestQualityInversion(test *testing.T) {
	data := []struct {
		input, expected Quality
	}{
		{Quality{PerfectType, 0}, Quality{PerfectType, 0}},
		{Quality{MajorType, 0}, Quality{MinorType, 0}},
		{Quality{MinorType, 0}, Quality{MajorType, 0}},
		{Quality{DiminishedType, 1}, Quality{AugmentedType, 1}},
		{Quality{AugmentedType, 1}, Quality{DiminishedType, 1}},
	}

	for i, t := range data {
		actual := t.input.Invert()
		if !actual.Eq(t.expected) {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}
