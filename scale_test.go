package musictheory

import (
	"reflect"
	"testing"
)

type scaleTest struct {
	root      Pitch
	intervals []Interval
	expected  []string
}

var scaleTests []scaleTest

func init() {
	scaleTests = []scaleTest{
		{NewPitch(C, Natural, 4), ChromaticIntervals, []string{"C4", "Db4", "D4", "Eb4", "E4", "F4", "Gb4", "G4", "Ab4", "A4", "Bb4", "B4"}},
		{NewPitch(C, Natural, 4), IonianIntervals, []string{"C4", "D4", "E4", "F4", "G4", "A4", "B4"}},
		{NewPitch(C, Natural, 4), DorianIntervals, []string{"C4", "D4", "Eb4", "F4", "G4", "A4", "Bb4"}},
		{NewPitch(C, Natural, 4), PhrygianIntervals, []string{"C4", "Db4", "Eb4", "F4", "G4", "Ab4", "Bb4"}},
		{NewPitch(C, Natural, 4), LydianIntervals, []string{"C4", "D4", "E4", "Gb4", "G4", "A4", "B4"}},
		{NewPitch(C, Natural, 4), MixolydianIntervals, []string{"C4", "D4", "E4", "F4", "G4", "A4", "Bb4"}},
		{NewPitch(C, Natural, 4), AeolianIntervals, []string{"C4", "D4", "Eb4", "F4", "G4", "Ab4", "Bb4"}},
		{NewPitch(C, Natural, 4), LocrianIntervals, []string{"C4", "Db4", "Eb4", "F4", "Gb4", "Ab4", "Bb4"}},
		{NewPitch(C, Natural, 4), MajorIntervals, []string{"C4", "D4", "E4", "F4", "G4", "A4", "B4"}},
		{NewPitch(C, Natural, 4), MinorIntervals, []string{"C4", "D4", "Eb4", "F4", "G4", "Ab4", "Bb4"}},
	}
}

func TestScales(test *testing.T) {
	for i, t := range scaleTests {
		scale := NewScale(t.root, t.intervals)
		actual := []string{}

		for _, p := range scale {
			actual = append(actual, p.(Pitch).Name(DescNames))
		}

		if !reflect.DeepEqual(actual, t.expected) {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}
