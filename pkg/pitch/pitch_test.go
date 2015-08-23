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
		actual := New(t.pitch, 0, Sharp).Name(SharpNames)
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
	{E, "Eb"},
	{F, "E"},
	{G, "Gb"},
	{A, "Ab"},
	{B, "Bb"},
}

func TestFlatPitchNames(test *testing.T) {
	for i, t := range flatNamesTests {
		actual := New(t.pitch, 0, Flat).Name(FlatNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

var doubleSharpNamesTests = []struct {
	pitch    int
	expected string
}{
	{C, "D"},
	{D, "E"},
	{E, "F#"},
	{F, "G"},
	{G, "A"},
	{A, "B"},
	{B, "C#"},
}

func TestDoubleSharpPitchNames(test *testing.T) {
	for i, t := range doubleSharpNamesTests {
		actual := New(t.pitch, 0, DoubleSharp).Name(SharpNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}

var doubleFlatNamesTests = []struct {
	pitch    int
	expected string
}{
	{C, "Bb"},
	{D, "C"},
	{E, "D"},
	{F, "Eb"},
	{G, "F"},
	{A, "G"},
	{B, "A"},
}

func TestDoubleFlatPitchNames(test *testing.T) {
	for i, t := range doubleFlatNamesTests {
		actual := New(t.pitch, 0, DoubleFlat).Name(FlatNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}
