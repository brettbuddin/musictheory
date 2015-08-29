package mt

import "math"

// Dotted makes a dotted duration
func Dotted(d Duration) Duration {
	return Duration{d.Value, 1}
}

// Durations
var (
	Longa               = Duration{0.25, 0}
	Breve               = Duration{0.5, 0}
	Whole               = Duration{1, 0}
	Half                = Duration{2, 0}
	Quarter             = Duration{4, 0}
	Eighth              = Duration{8, 0}
	Sixteenth           = Duration{16, 0}
	ThirtySecond        = Duration{32, 0}
	SixtyFourth         = Duration{64, 0}
	HundredTwentyEighth = Duration{128, 0}
)

// Duration represents a note's duration
type Duration struct {
	Value float64
	Dots  int
}

// Seconds returns the time in seconds the note's duration lasts.
// Calculated based on what unit gets the beat and what the BPM is.
func (d Duration) Seconds(unit Duration, bpm int) float64 {
	unitDuration := math.Ceil(unit.Value)
	val := (60.0 / float64(bpm)) / (float64(d.Value) / 4.0) / (float64(unitDuration) / 4.0)
	return float64(val*2) - val/math.Pow(2, float64(d.Dots))
}

func (d Duration) String() string {
	switch d.Value {
	case 0.25:
		return "longa"
	case 0.5:
		return "breve"
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
