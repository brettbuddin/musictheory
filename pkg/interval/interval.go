package interval

import (
	"fmt"
	mt_math "github.com/brettbuddin/mt/pkg/math"
)

// Quality types
const (
	PerfectT QualityType = iota
	MajorT
	MinorT
	AugmentedT
	DiminishedT
)

// Intervals
var (
	Perfect          = qualityInterval(Quality{PerfectT, 0})
	Major            = qualityInterval(Quality{MajorT, 0})
	Minor            = qualityInterval(Quality{MinorT, 0})
	Augmented        = qualityInterval(Quality{AugmentedT, 1})
	DoublyAugmented  = qualityInterval(Quality{AugmentedT, 2})
	Diminished       = qualityInterval(Quality{DiminishedT, 1})
	DoublyDiminished = qualityInterval(Quality{DiminishedT, 2})
	Octave           = Interval{1, 0, 0}
)

func qualityInterval(quality Quality) func(int) Interval {
	return func(step int) Interval {
		diatonic := normalizeDiatonic(step - 1)
		diff := qualityDiff(quality, canBePerfect(diatonic))
		octaves := diatonicOctaves(step - 1)
		return New(step, octaves, diff)
	}
}

// New Interval
func New(step, octaves, offset int) Interval {
	diatonic := normalizeDiatonic(step - 1)
	chromatic := DiatonicToChromatic(diatonic) + offset

	return Interval{octaves, diatonic, chromatic}
}

// Interval represents an interval in 12-tone equal temperament
type Interval struct {
	octaves   int
	diatonic  int
	chromatic int
}

func (i Interval) String() string {
	return fmt.Sprintf("(octaves: %d, diatonic: %d, chromatic: %d)", i.octaves, i.diatonic, i.chromatic)
}

// Octaves returns the octave component
func (i Interval) Octaves() int {
	return i.octaves
}

// Diff returns the difference between the chromatic component and the chromatized diatonic
func (i Interval) ChromaticDiff() int {
	return i.chromatic - DiatonicToChromatic(i.diatonic)
}

// Diatonic returns the diatonic component
func (i Interval) Diatonic() int {
	return i.diatonic
}

// Chromatic returns the chromatic component
func (i Interval) Chromatic() int {
	return i.chromatic
}

// Semitones returns the total number of semitones that make up the interval
func (i Interval) Semitones() int {
	return i.octaves*12 + i.chromatic
}

// Quality returns the Quality
func (i Interval) Quality() Quality {
	quality := diffQuality(i.Chromatic()-DiatonicToChromatic(i.Diatonic()), canBePerfect(i.Diatonic()))

	if i.Octaves() < 0 {
		return quality.Invert()
	}

	return quality
}

// Transpose returns a new Interval that has been transposed by the given Interval
func (i Interval) Transpose(o Interval) Interval {
	diatonic := i.Diatonic() + o.Diatonic()
	diatonicOctaves := diatonicOctaves(diatonic)
	diatonicRemainder := normalizeDiatonic(diatonic)

	octaves := i.Octaves() + o.Octaves() + diatonicOctaves
	chromatic := normalizeChromatic(i.Chromatic() + o.Chromatic())

	return Interval{octaves, diatonicRemainder, chromatic}
}

// Negate returns a new, negated Interval
func (i Interval) Negate() Interval {
	if i.diatonic == 0 && i.chromatic == 0 {
		return Interval{-i.octaves, i.diatonic, i.chromatic}
	}

	return Interval{-(i.octaves + 1), inverseDiatonic(i.diatonic), inverseChromatic(i.chromatic)}
}

// QualityType represents the type a Quality can take
type QualityType int

func (q QualityType) String() string {
	switch q {
	case PerfectT:
		return "perfect"
	case MajorT:
		return "major"
	case MinorT:
		return "minor"
	case AugmentedT:
		return "augmented"
	case DiminishedT:
		return "diminished"
	default:
		return "unknown"
	}
}

// Quality describes the quality of an interval
type Quality struct {
	Type QualityType
	Size int
}

// Invert returns a new, inverted Quality
func (q Quality) Invert() Quality {
	switch q.Type {
	case PerfectT:
		return q
	case MajorT:
		return Quality{MinorT, q.Size}
	case MinorT:
		return Quality{MajorT, q.Size}
	case AugmentedT:
		return Quality{DiminishedT, q.Size}
	case DiminishedT:
		return Quality{AugmentedT, q.Size}
	default:
		panic(fmt.Sprintf("invalid type: %d", q.Type))
	}
}

// Eq checks two Qualities for equality
func (q Quality) Eq(o Quality) bool {
	return q.Type == o.Type && q.Size == o.Size
}

// DiatonicToChromatic converts a diatonic value to the chromatic equivalent
func DiatonicToChromatic(interval int) int {
	if interval >= len(diatonicToChromaticLookup) {
		panic(fmt.Sprintf("interval out of range: %d", interval))
	}

	return diatonicToChromaticLookup[interval]
}

var diatonicToChromaticLookup = []int{0, 2, 4, 5, 7, 9, 11}

func qualityDiff(q Quality, perfect bool) int {
	if q.Type == PerfectT || q.Type == MajorT {
		return 0
	} else if q.Type == MinorT {
		return -1
	} else if q.Type == AugmentedT {
		return q.Size
	} else if q.Type == DiminishedT {
		if perfect {
			return -q.Size
		}
		return -(q.Size + 1)
	}
	panic("invalid quality")
}

func diffQuality(diff int, perfect bool) Quality {
	if perfect {
		if diff == 0 {
			return Quality{PerfectT, 0}
		} else if diff > 0 {
			return Quality{AugmentedT, diff}
		}

		return Quality{DiminishedT, -diff}
	}

	if diff == 0 {
		return Quality{MajorT, 0}
	} else if diff == -1 {
		return Quality{MinorT, 0}
	} else if diff > 0 {
		return Quality{AugmentedT, diff}
	}

	return Quality{DiminishedT, -(diff + 1)}
}

func canBePerfect(interval int) bool {
	return interval == 0 || interval == 3 || interval == 4
}

func normalizeChromatic(v int) int {
	return int(mt_math.Mod(float64(v), 12))
}

func normalizeDiatonic(v int) int {
	return int(mt_math.Mod(float64(v), 7))
}

func diatonicOctaves(v int) int {
	return v / 7
}

func inverseChromatic(v int) int {
	return 12 - v
}

func inverseDiatonic(v int) int {
	return 7 - v
}
