package mt

import (
	"fmt"
	"math"
)

// Modifier offset values
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
	modifierNames  = [5]string{"bb", "b", "", "#", "x"}
	pitchNames     = [7]string{"C", "D", "E", "F", "G", "A", "B"}
	namesForFlats  = [12]int{0, 1, 1, 2, 2, 3, 4, 4, 5, 5, 6, 6}
	namesForSharps = [12]int{0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6}
	semitone       = math.Pow(2, 1.0/12.0)
	middleA        = NewPitch(A, Natural, 4)
)

// FlatNames maps an modifier to a correspending diatonic as flats
func FlatNames(i int) int {
	return namesForFlats[normalizeChromatic(i)]
}

// SharpNames maps an modifier to a correspending diatonic as sharps
func SharpNames(i int) int {
	return namesForSharps[normalizeChromatic(i)]
}

// NewPitch builds a new Pitch
func NewPitch(diatonic, modifier, octaves int) Pitch {
	return Pitch{NewInterval(diatonic, octaves, modifier)}
}

// Pitch represents an absolute pitch in 12-tone equal temperament
type Pitch struct {
	Interval
}

type ModifierStrategy func(int) int

// Name returns the name of the pitch using a particular name strategy (either SharpNames or FlatNames). The result is
// in scientific pitch notation format.
func (p Pitch) Name(strategy ModifierStrategy) string {
	semitones := normalizeChromatic(p.Chromatic)
	nameIndex := strategy(semitones)
	delta := semitones - diatonicToChromatic(nameIndex)

	if delta == 0 {
		return fmt.Sprintf("%s%d", pitchNames[nameIndex], p.Octaves)
	}
	return fmt.Sprintf("%s%s%d", pitchNames[nameIndex], modifierName(delta+2), p.Octaves)
}

// Transpose transposes a pitch by a given interval
func (p Pitch) Transpose(i Interval) Pitch {
	return Pitch{p.Interval.Transpose(i)}
}

// Eq determines if another pitch is the same
func (p Pitch) Eq(o Pitch) bool {
	return p.Interval.Eq(o.Interval)
}

// Freq returns the absolute frequency of a pitch in Hz
func (p Pitch) Freq() float64 {
	return concertFrequency * math.Pow(semitone, float64(p.Semitones()-middleA.Semitones()))
}

// MIDI returns the MIDI note number of the pitch
func (p Pitch) MIDI() int {
	return p.Semitones() + 24
}

func modifierName(i int) string {
	return modifierNames[int(mod(float64(i), float64(len(modifierNames))))]
}
