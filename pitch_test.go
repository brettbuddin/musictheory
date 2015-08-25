package mt

import (
	"testing"
)

var sharpNamesTests = []struct {
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

func TestSharpPitchNames(test *testing.T) {
	for i, t := range sharpNamesTests {
		actual := NewPitch(t.pitch, 0, Sharp).Name(SharpNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

var flatNamesTests = []struct {
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

func TestFlatPitchNames(test *testing.T) {
	for i, t := range flatNamesTests {
		actual := NewPitch(t.pitch, 0, Flat).Name(FlatNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

var doubleSharpNamesTests = []struct {
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

func TestDoubleSharpPitchNames(test *testing.T) {
	for i, t := range doubleSharpNamesTests {
		actual := NewPitch(t.pitch, 0, DoubleSharp).Name(SharpNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

var doubleFlatNamesTests = []struct {
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

func TestDoubleFlatPitchNames(test *testing.T) {
	for i, t := range doubleFlatNamesTests {
		actual := NewPitch(t.pitch, 0, DoubleFlat).Name(FlatNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}
