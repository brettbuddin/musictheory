package musictheory

import (
	"testing"
)

func TestParsePitch(test *testing.T) {
	data := []struct {
		input    string
		expected Pitch
	}{
		{"C4", NewPitch(C, Natural, 4)},
		{"C#4", NewPitch(C, Sharp, 4)},
		{"Ab3", NewPitch(A, Flat, 3)},
		{"Abb3", NewPitch(A, DoubleFlat, 3)},
		{"Cx4", NewPitch(C, DoubleSharp, 4)},
	}

	for i, t := range data {
		actual, err := ParsePitch(t.input)
		if err != nil {
			test.Error(err)
		}

		if !actual.Eq(t.expected) {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}
