package pitch

import (
	"testing"
)

var sharpNamesTests = []struct {
	pitch    int
	expected string
}{
	{C, "C#"},
	{D, "D#"},
	{E, "F"},
	{F, "F#"},
	{G, "G#"},
	{A, "A#"},
	{B, "C"},
}

func TestSharpPitchNames(test *testing.T) {
	for i, t := range sharpNamesTests {
		actual := New(t.pitch, Sharp).Name(SharpNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

var flatNamesTests = []struct {
	pitch    int
	expected string
}{
	{C, "B"},
	{D, "Db"},
	{F, "E"},
	{E, "Eb"},
	{A, "Ab"},
	{B, "Bb"},
}

func TestFlatPitchNames(test *testing.T) {
	for i, t := range flatNamesTests {
		actual := New(t.pitch, Flat).Name(FlatNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}
