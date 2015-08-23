package pitch

import (
	"fmt"
	"github.com/brettbuddin/mt/pkg/interval"
	mt_math "github.com/brettbuddin/mt/pkg/math"
)

const (
	Semitone = 1
	Tone     = 2
	Ditone   = 3
	Tritone  = 6

	DoubleFlat  = -2
	Flat        = -1
	Natural     = 0
	Sharp       = 1
	DoubleSharp = 2
)

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
	AccidentalNames = [5]string{"bb", "b", "", "#", "x"}
	PitchNames      = [7]string{"C", "D", "E", "F", "G", "A", "B"}
	namesForFlats   = [12]int{0, 1, 1, 2, 2, 3, 4, 4, 5, 5, 6, 6}
	namesForSharps  = [12]int{0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6}

	FlatNames  = NameStrategyFunc(func(i int) int { return namesForFlats[int(mt_math.Mod(float64(i), 12))] })
	SharpNames = NameStrategyFunc(func(i int) int { return namesForSharps[int(mt_math.Mod(float64(i), 12))] })
)

type NameStrategy interface {
	GetMappedIndex(int) int
}

type NameStrategyFunc func(int) int

func (f NameStrategyFunc) GetMappedIndex(i int) int {
	return f(i)
}

func New(semitone int, accidental int) Pitch {
	return Pitch{interval.New(semitone, accidental)}
}

type Pitch struct {
	interval interval.Interval
}

func (p Pitch) Name(s NameStrategy) string {
	nameIndex := s.GetMappedIndex(p.interval.Chroma())
	diff := p.interval.Diff()
	accIndex := diff + 2
	diatonic := p.interval.Diatonic()

	// TODO: Find a better way to detect natural semitones.

	if (diatonic == 0 || diatonic == 3) && diff == -1 {
		return PitchNames[nameIndex]
	}

	if (diatonic == 2 || diatonic == 6) && diff == 1 {
		return PitchNames[nameIndex]
	}

	return fmt.Sprintf("%s%s", PitchNames[nameIndex], AccidentalNames[accIndex])
}
