package pitch

import (
	"fmt"
	"github.com/brettbuddin/mt/pkg/interval"
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
	MiddleOctave        = 4
	UseFancyAccidentals = false

	accidentalNames      = [5]string{"bb", "b", "", "#", "x"}
	fancyAccidentalNames = [5]string{"♭♭", "♭", "", "♯", "𝄪"}
	pitchNames           = [7]string{"C", "D", "E", "F", "G", "A", "B"}
	namesForFlats        = [12]int{0, 1, 1, 2, 2, 3, 4, 4, 5, 5, 6, 6}
	namesForSharps       = [12]int{0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6}
	semitone             = math.Pow(2, 1.0/12.0)
	middleA              = NewAbsolute(A, MiddleOctave, Natural)
)

// FlatNames maps an accidental to a correspending diatonic as flats
func FlatNames(i int) int {
	return namesForFlats[int(mt_math.Mod(float64(i), 12))]
}

// SharpNames maps an accidental to a correspending diatonic as sharps
func SharpNames(i int) int {
	return namesForSharps[int(mt_math.Mod(float64(i), 12))]
}

type nameStrategy interface {
	GetMappedIndex(int) int
}

type nameStrategyFunc func(int) int

func (f nameStrategyFunc) GetMappedIndex(i int) int {
	return f(i)
}

// New builds a new Pitch at an octave relative to MiddleOctave
func New(diatonic, octaves, accidental int) Pitch {
	return Pitch{interval.New(diatonic, MiddleOctave+octaves, accidental)}
}

// NewAbsolute builds a new Pitch at an absolute octave
func NewAbsolute(diatonic, octave, accidental int) Pitch {
	return Pitch{interval.New(diatonic, octave, accidental)}
}

// Pitch represents an absolute pitch in 12-tone equal temperament
type Pitch struct {
	interval.Interval
}

type AccidentalStrategy func(int) int

// Name returns the name of the pitch using a particular name strategy (either SharpNames or FlatNames). The result is
// in scientific pitch notation format.
func (p Pitch) Name(strategy AccidentalStrategy) string {
	semitones := int(mt_math.Mod(float64(p.Semitones()), 12.0))
	nameIndex := strategy(semitones)
	delta := semitones - interval.DiatonicToChromatic(nameIndex)

	if delta == 0 {
		return fmt.Sprintf("%s%d", pitchNames[nameIndex], p.Octaves())
	}
	return fmt.Sprintf("%s%s%d", pitchNames[nameIndex], accidentalName(delta+2), p.Octaves())
}

// Transpose transposes a pitch by a given interval
func (p Pitch) Transpose(i interval.Interval) Pitch {
	return Pitch{p.Interval.Transpose(i)}
}

// Freq returns the absolute frequency of a pitch in Hz
func (p Pitch) Freq() float64 {
	return concertFrequency * math.Pow(semitone, float64(p.Semitones()-middleA.Semitones()))
}

// MIDI returns the MIDI note number of the pitch
func (p Pitch) MIDI() int {
	return p.Semitones() + 12
}

func accidentalName(i int) string {
	if UseFancyAccidentals {
		return fancyAccidentalNames[int(mt_math.Mod(float64(i), float64(len(accidentalNames))))]
	}
	return accidentalNames[int(mt_math.Mod(float64(i), float64(len(accidentalNames))))]
}
