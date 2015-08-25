package mt

import (
	"fmt"
	mt_math "github.com/brettbuddin/mt/pkg/math"
	"math"
)

// Accidental offset values
const (
	DoubleFlat  = -2
	Flat        = -1
	Natural     = 0
	Sharp       = 1
	DoubleSharp = 2

	concertFrequency = 440.0
)

// Note steps
const (
	C int = iota + 1
	D
	E
	F
	G
	A
	B
)

var (
	MiddleOctave = 4

	accidentalNames = [5]string{"bb", "b", "", "#", "x"}
	pitchNames      = [7]string{"C", "D", "E", "F", "G", "A", "B"}
	namesForFlats   = [12]int{0, 1, 1, 2, 2, 3, 4, 4, 5, 5, 6, 6}
	namesForSharps  = [12]int{0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6}
	semitone        = math.Pow(2, 1.0/12.0)
	middleA         = NewPitch(A, 0, Natural)
)

// FlatNames maps an accidental to a correspending diatonic as flats
func FlatNames(i int) int {
	return namesForFlats[normalizeChromatic(i)]
}

// SharpNames maps an accidental to a correspending diatonic as sharps
func SharpNames(i int) int {
	return namesForSharps[normalizeChromatic(i)]
}

// NewPitch builds a new Pitch at an octave relative to MiddleOctave
func NewPitch(diatonic, octaves, accidental int) Pitch {
	return Pitch{NewInterval(diatonic, MiddleOctave+octaves, accidental)}
}

// Pitch represents an absolute pitch in 12-tone equal temperament
type Pitch struct {
	Interval
}

type AccidentalStrategy func(int) int

// Name returns the name of the pitch using a particular name strategy (either SharpNames or FlatNames). The result is
// in scientific pitch notation format.
func (p Pitch) Name(strategy AccidentalStrategy) string {
	semitones := normalizeChromatic(p.Semitones())
	nameIndex := strategy(semitones)
	delta := semitones - diatonicToChromatic(nameIndex)

	if delta == 0 {
		return fmt.Sprintf("%s%d", pitchNames[nameIndex], p.Octaves())
	}
	return fmt.Sprintf("%s%s%d", pitchNames[nameIndex], accidentalName(delta+2), p.Octaves())
}

// Transpose transposes a pitch by a given interval
func (p Pitch) Transpose(i Interval) Pitch {
	return Pitch{p.Interval.Transpose(i)}
}

// Freq returns the absolute frequency of a pitch in Hz
func (p Pitch) Freq() float64 {
	return concertFrequency * math.Pow(semitone, float64(p.Semitones()-middleA.Semitones()))
}

// MIDI returns the MIDI note number of the pitch
func (p Pitch) MIDI() int {
	return p.Semitones() + 24
}

func accidentalName(i int) string {
	return accidentalNames[int(mt_math.Mod(float64(i), float64(len(accidentalNames))))]
}
