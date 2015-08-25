package mt

import (
	"reflect"
	"testing"
)

type scaleTest struct {
	root      Pitch
	intervals Intervals
	expected  []string
}

var scaleTests []scaleTest

func init() {
	scaleTests = []scaleTest{
		{NewPitch(C, 0, Natural), ChromaticIntervals, []string{"C4", "Db4", "D4", "Eb4", "E4", "F4", "Gb4", "G4", "Ab4", "A4", "Bb4", "B4"}},
		{NewPitch(C, 0, Natural), IonianIntervals, []string{"C4", "D4", "E4", "F4", "G4", "A4", "B4"}},
		{NewPitch(C, 0, Natural), DorianIntervals, []string{"C4", "D4", "Eb4", "F4", "G4", "A4", "Bb4"}},
		{NewPitch(C, 0, Natural), PhrygianIntervals, []string{"C4", "Db4", "Eb4", "F4", "G4", "Ab4", "Bb4"}},
		{NewPitch(C, 0, Natural), LydianIntervals, []string{"C4", "D4", "E4", "Gb4", "G4", "A4", "B4"}},
		{NewPitch(C, 0, Natural), MixolydianIntervals, []string{"C4", "D4", "E4", "F4", "G4", "A4", "Bb4"}},
		{NewPitch(C, 0, Natural), AeolianIntervals, []string{"C4", "D4", "Eb4", "F4", "G4", "Ab4", "Bb4"}},
		{NewPitch(C, 0, Natural), LocrianIntervals, []string{"C4", "Db4", "Eb4", "F4", "Gb4", "Ab4", "Bb4"}},
		{NewPitch(C, 0, Natural), MajorIntervals, []string{"C4", "D4", "E4", "F4", "G4", "A4", "B4"}},
		{NewPitch(C, 0, Natural), MinorIntervals, []string{"C4", "D4", "Eb4", "F4", "G4", "Ab4", "Bb4"}},
	}
}

func TestScales(test *testing.T) {
	for i, t := range scaleTests {
		scale := NewScale(t.root, t.intervals)
		actual := []string{}

		for _, p := range scale {
			actual = append(actual, p.Name(FlatNames))
		}

		if !reflect.DeepEqual(actual, t.expected) {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}
