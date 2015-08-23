package pitch

import (
	"testing"
)

var sharpNamesTests = []struct {
	pitch    int
	expected string
}{
	{C, "C#0"},
	{D, "D#0"},
	{E, "F0"},
	{F, "F#0"},
	{G, "G#0"},
	{A, "A#0"},
	{B, "C0"},
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
	{C, "B0"},
	{D, "Db0"},
	{E, "Eb0"},
	{F, "E0"},
	{G, "Gb0"},
	{A, "Ab0"},
	{B, "Bb0"},
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
	{C, "D0"},
	{D, "E0"},
	{E, "F#0"},
	{F, "G0"},
	{G, "A0"},
	{A, "B0"},
	{B, "C#0"},
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
	{C, "Bb0"},
	{D, "C0"},
	{E, "D0"},
	{F, "Eb0"},
	{G, "F0"},
	{A, "G0"},
	{B, "A0"},
}

func TestDoubleFlatPitchNames(test *testing.T) {
	for i, t := range doubleFlatNamesTests {
		actual := New(t.pitch, 0, DoubleFlat).Name(FlatNames)
		if actual != t.expected {
			test.Errorf("index=%d actual=%s expected=%s", i, actual, t.expected)
		}
	}
}
