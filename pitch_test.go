package mt

import (
	"math"
	"testing"
)

const tolerance = 0.000001

func TestSharpPitchNames(test *testing.T) {
	data := []struct {
		pitch    int
		expected string
	}{
		{C, "C#4"},
		{D, "D#4"},
		{E, "F4"},
		{F, "F#4"},
		{G, "G#4"},
		{A, "A#4"},
		{B, "C4"},
	}

	for i, t := range data {
		actual := NewPitch(t.pitch, 0, Sharp).Name(SharpNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

func TestFlatPitchNames(test *testing.T) {
	data := []struct {
		pitch    int
		expected string
	}{
		{C, "B4"},
		{D, "Db4"},
		{E, "Eb4"},
		{F, "E4"},
		{G, "Gb4"},
		{A, "Ab4"},
		{B, "Bb4"},
	}

	for i, t := range data {
		actual := NewPitch(t.pitch, 0, Flat).Name(FlatNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

func TestDoubleSharpPitchNames(test *testing.T) {
	data := []struct {
		pitch    int
		expected string
	}{
		{C, "D4"},
		{D, "E4"},
		{E, "F#4"},
		{F, "G4"},
		{G, "A4"},
		{A, "B4"},
		{B, "C#4"},
	}

	for i, t := range data {
		actual := NewPitch(t.pitch, 0, DoubleSharp).Name(SharpNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

func TestDoubleFlatPitchNames(test *testing.T) {
	data := []struct {
		pitch    int
		expected string
	}{
		{C, "Bb4"},
		{D, "C4"},
		{E, "D4"},
		{F, "Eb4"},
		{G, "F4"},
		{A, "G4"},
		{B, "A4"},
	}

	for i, t := range data {
		actual := NewPitch(t.pitch, 0, DoubleFlat).Name(FlatNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

func TestFrequency(test *testing.T) {
	data := []struct {
		input    Pitch
		expected float64
	}{
		{NewPitch(C, 0, Natural), 261.625565},
		{NewPitch(A, 0, Natural), 440.0},
	}

	for _, t := range data {
		actual := t.input.Freq()

		if closeEqualFloat64(actual, t.expected) {
			test.Errorf("input=%s output=%f, expected=%f",
				t.input,
				actual,
				t.expected)
		}
	}
}

func TestMIDI(test *testing.T) {
	data := []struct {
		input    Pitch
		expected int
	}{
		{NewPitch(C, -1, Natural), 60},
		{NewPitch(C, 0, Natural), 72},
		{NewPitch(A, 0, Natural), 81},
	}

	for _, t := range data {
		actual := t.input.MIDI()

		if actual != t.expected {
			test.Errorf("input=%s output=%f, expected=%f",
				t.input,
				actual,
				t.expected)
		}
	}
}

func closeEqualFloat64(actual, expected float64) bool {
	return math.Abs(actual-expected) >= tolerance
}
