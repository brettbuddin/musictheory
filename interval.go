package mt

import (
	"fmt"
	"math"
)

// Quality types
const (
	PerfectType QualityType = iota
	MajorType
	MinorType
	AugmentedType
	DiminishedType
)

// Specific intervals

var Octave = Interval{1, 0, 0}

type IntervalFunc func(int) Interval

func Perfect(step int) Interval {
	return qualityInterval(step, Quality{PerfectType, 0})
}

func Major(step int) Interval {
	return qualityInterval(step, Quality{MajorType, 0})
}

func Minor(step int) Interval {
	return qualityInterval(step, Quality{MinorType, 0})
}

func Augmented(step int) Interval {
	return qualityInterval(step, Quality{AugmentedType, 1})
}

func DoublyAugmented(step int) Interval {
	return qualityInterval(step, Quality{AugmentedType, 2})
}

func Diminished(step int) Interval {
	return qualityInterval(step, Quality{DiminishedType, 1})
}

func DoublyDiminished(step int) Interval {
	return qualityInterval(step, Quality{DiminishedType, 2})
}

func qualityInterval(step int, quality Quality) Interval {
	diatonic := normalizeDiatonic(step - 1)
	diff := qualityDiff(quality, canBePerfect(diatonic))
	octaves := diatonicOctaves(step - 1)
	return NewInterval(step, octaves, diff)
}

// NewInterval builds a new Interval
func NewInterval(step, octaves, offset int) Interval {
	diatonic := normalizeDiatonic(step - 1)
	chromatic := diatonicToChromatic(diatonic) + offset

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
	return i.chromatic - diatonicToChromatic(i.diatonic)
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
	quality := diffQuality(i.Chromatic()-diatonicToChromatic(i.Diatonic()), canBePerfect(i.Diatonic()))

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
	case PerfectType:
		return "perfect"
	case MajorType:
		return "major"
	case MinorType:
		return "minor"
	case AugmentedType:
		return "augmented"
	case DiminishedType:
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
	case PerfectType:
		return q
	case MajorType:
		return Quality{MinorType, q.Size}
	case MinorType:
		return Quality{MajorType, q.Size}
	case AugmentedType:
		return Quality{DiminishedType, q.Size}
	case DiminishedType:
		return Quality{AugmentedType, q.Size}
	default:
		panic(fmt.Sprintf("invalid type: %d", q.Type))
	}
}

// Eq checks two Qualities for equality
func (q Quality) Eq(o Quality) bool {
	return q.Type == o.Type && q.Size == o.Size
}

func (q Quality) String() string {
	switch q.Type {
	case PerfectType, MajorType, MinorType:
		return fmt.Sprintf("%s", q.Type)
	case AugmentedType, DiminishedType:
		return fmt.Sprintf("%s(%d)", q.Type, q.Size)
	default:
		return "unknown"
	}
}

func diatonicToChromatic(interval int) int {
	if interval >= len(diatonicToChromaticLookup) {
		panic(fmt.Sprintf("interval out of range: %d", interval))
	}

	return diatonicToChromaticLookup[interval]
}

var diatonicToChromaticLookup = []int{0, 2, 4, 5, 7, 9, 11}

func qualityDiff(q Quality, perfect bool) int {
	if q.Type == PerfectType || q.Type == MajorType {
		return 0
	} else if q.Type == MinorType {
		return -1
	} else if q.Type == AugmentedType {
		return q.Size
	} else if q.Type == DiminishedType {
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
			return Quality{PerfectType, 0}
		} else if diff > 0 {
			return Quality{AugmentedType, diff}
		}

		return Quality{DiminishedType, -diff}
	}

	if diff == 0 {
		return Quality{MajorType, 0}
	} else if diff == -1 {
		return Quality{MinorType, 0}
	} else if diff > 0 {
		return Quality{AugmentedType, diff}
	}

	return Quality{DiminishedType, -(diff + 1)}
}

func canBePerfect(interval int) bool {
	return interval == 0 || interval == 3 || interval == 4
}

func normalizeChromatic(v int) int {
	return int(mod(float64(v), 12))
}

func normalizeDiatonic(v int) int {
	return int(mod(float64(v), 7))
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

func mod(n, m float64) float64 {
	out := math.Mod(n, m)
	if out < 0 {
		out += m
	}
	return out
}
