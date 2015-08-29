package mt

import (
	//"math"
	"time"
)

// Dotted makes a dotted duration
func Dotted(d Duration) Duration {
	return Duration{d.Value, true, false}
}

// Triplet makes a triplet duration
func Triplet(d Duration) Duration {
	return Duration{d.Value, false, true}
}

const ns = 1000000000

// Durations
var (
	D1   = Duration{1, false, false}
	D2   = Duration{2, false, false}
	D4   = Duration{4, false, false}
	D8   = Duration{8, false, false}
	D16  = Duration{16, false, false}
	D32  = Duration{32, false, false}
	D64  = Duration{64, false, false}
	D128 = Duration{128, false, false}
)

// Duration represents a note's duration
type Duration struct {
	Value   int
	Dot     bool
	Triplet bool
}

// Seconds returns the time in nanoseconds the note's duration lasts.
// Calculated based on what unit gets the beat and what the BPM is.
func (d Duration) Time(unit Duration, bpm int) time.Duration {
	val := (60.0 / float64(bpm)) / (float64(d.Value) / 4.0) / (float64(unit.Value) / 4.0)

	if d.Dot {
		val += val / 2.0
	}

	if d.Triplet {
		val = val / 3.0
	}

	return time.Duration(val * ns)
}

func (d Duration) String() string {
	switch d.Value {
	case 1:
		return "whole"
	case 2:
		return "half"
	case 4:
		return "quarter"
	case 8:
		return "eighth"
	case 16:
		return "sixteenth"
	case 32:
		return "thirty-second"
	case 64:
		return "sixty-fourth"
	case 128:
		return "hundred-twenty-eighth"
	default:
		return "unknown"
	}
}

type Note struct {
	Pitch
	Duration
}

// NewNote creates a new note
func NewNote(diatonic, octaves, accidental int, duration Duration) Note {
	return Note{NewPitch(diatonic, octaves, accidental), duration}
}

// Transpose transposes a note by a given interval
func (n Note) Transpose(i Interval) Note {
	return Note{n.Pitch.Transpose(i), n.Duration}
}
