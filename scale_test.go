package musictheory

import (
	"reflect"
	"testing"
)

type scaleTest struct {
	root      Pitch
	intervals []Interval
	octaves   int
	expected  []string
}

var scaleTests []scaleTest

func init() {
	scaleTests = []scaleTest{
		{NewPitch(C, Natural, 4), ChromaticIntervals, 1, []string{"C4", "Db4", "D4", "Eb4", "E4", "F4", "Gb4", "G4", "Ab4", "A4", "Bb4", "B4"}},
		{NewPitch(C, Natural, 4), IonianIntervals, 1, []string{"C4", "D4", "E4", "F4", "G4", "A4", "B4"}},
		{NewPitch(C, Natural, 4), DorianIntervals, 1, []string{"C4", "D4", "Eb4", "F4", "G4", "A4", "Bb4"}},
		{NewPitch(C, Natural, 4), PhrygianIntervals, 1, []string{"C4", "Db4", "Eb4", "F4", "G4", "Ab4", "Bb4"}},
		{NewPitch(C, Natural, 4), LydianIntervals, 1, []string{"C4", "D4", "E4", "Gb4", "G4", "A4", "B4"}},
		{NewPitch(C, Natural, 4), MixolydianIntervals, 1, []string{"C4", "D4", "E4", "F4", "G4", "A4", "Bb4"}},
		{NewPitch(C, Natural, 4), AeolianIntervals, 1, []string{"C4", "D4", "Eb4", "F4", "G4", "Ab4", "Bb4"}},
		{NewPitch(C, Natural, 4), LocrianIntervals, 1, []string{"C4", "Db4", "Eb4", "F4", "Gb4", "Ab4", "Bb4"}},
		{NewPitch(C, Natural, 4), MajorIntervals, 2, []string{"C4", "D4", "E4", "F4", "G4", "A4", "B4", "C5", "D5", "E5", "F5", "G5", "A5", "B5"}},
		{NewPitch(C, Natural, 4), MinorIntervals, 1, []string{"C4", "D4", "Eb4", "F4", "G4", "Ab4", "Bb4"}},
		{NewPitch(E, Natural, 4), MinorPentatonicIntervals, 1, []string{"E4", "G4", "A4", "B4", "D5"}},
		{NewPitch(E, Flat, 4), MajorPentatonicIntervals, 1, []string{"Eb4", "F4", "G4", "Bb4", "C5"}},
	}
}

func TestScales(test *testing.T) {
	for i, t := range scaleTests {
		scale := NewScale(t.root, t.intervals, t.octaves)
		actual := []string{}

		for _, p := range scale {
			actual = append(actual, p.(Pitch).Name(DescNames))
		}

		if !reflect.DeepEqual(actual, t.expected) {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}
