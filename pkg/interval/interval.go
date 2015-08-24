package interval

import (
	"fmt"
	mt_math "github.com/brettbuddin/mt/pkg/math"
)

const (
	PerfectT int = iota
	MajorT
	MinorT
	AugmentedT
	DiminishedT
)

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
	return func(val int) Interval {
		diatonic := int(mt_math.Mod(float64(val-1), 7))
		diff := qualityDiff(perfect(diatonic), quality)
		octaves := int((val - 1) / 7.0)
		return New(val, octaves, diff)
	}
}

func New(diatonic, octaves, offset int) Interval {
	diatonic = int(mt_math.Mod(float64(diatonic-1), 7))
	chromatic := DiatonicToChromatic(diatonic) + offset

	return Interval{octaves, diatonic, chromatic}
}

type Interval struct {
	octaves   int
	diatonic  int
	chromatic int
}

func (i Interval) String() string {
	return fmt.Sprintf("(octaves: %d, diatonic: %d, chromatic: %d)", i.octaves, i.diatonic, i.chromatic)
}

func (i Interval) Octaves() int {
	return i.octaves
}

func (i Interval) Diff() int {
	return i.chromatic - DiatonicToChromatic(i.diatonic)
}

func (i Interval) Diatonic() int {
	return i.diatonic
}

func (i Interval) Chromatic() int {
	return i.chromatic
}

func (i Interval) Semitones() int {
	return i.octaves*12 + i.chromatic
}

func (i Interval) Quality() Quality {
	quality := diffQuality(perfect(i.Diatonic()), i.Chromatic()-DiatonicToChromatic(i.Diatonic()))

	if i.Octaves() < 0 {
		return quality.Invert()
	}

	return quality
}

func (i Interval) HasQualityType(t int) bool {
	return i.Quality().Type == t
}

func (i Interval) AddInterval(o Interval) Interval {
	diatonics := i.Diatonic() + o.Diatonic()
	diatonicOctaves := diatonics / 7.0
	diatonicRemainder := int(mt_math.Mod(float64(diatonics), 7.0))

	octaves := i.Octaves() + o.Octaves() + diatonicOctaves
	chromatic := i.Chromatic() + o.Chromatic()
	chromatic = int(mt_math.Mod(float64(chromatic), 12.0))

	return Interval{octaves, diatonicRemainder, chromatic}
}

func (i Interval) Negate() Interval {
	if i.diatonic == 0 && i.chromatic == 0 {
		return Interval{-i.octaves, i.diatonic, i.chromatic}
	}

	return Interval{-(i.octaves + 1), 7 - i.diatonic, 12 - i.chromatic}
}

type Quality struct {
	Type, Size int
}

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

func (q Quality) Eq(o Quality) bool {
	return q.Type == o.Type && q.Size == o.Size
}

func DiatonicToChromatic(interval int) int {
	if interval >= len(diatonicToChromaticLookup) {
		panic(fmt.Sprintf("interval out of range: %d", interval))
	}

	return diatonicToChromaticLookup[interval]
}

var diatonicToChromaticLookup = []int{0, 2, 4, 5, 7, 9, 11}

func qualityDiff(perfect bool, q Quality) int {
	if q.Type == PerfectT || q.Type == MajorT {
		return 0
	} else if q.Type == MinorT {
		return -1
	} else if q.Type == AugmentedT {
		return q.Size
	} else if q.Type == DiminishedT {
		if perfect {
			return -q.Size
		} else {
			return -(q.Size + 1)
		}
	}
	panic("invalid quality")
}

func diffQuality(perfect bool, diff int) Quality {
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

func perfect(interval int) bool {
	return interval == 0 || interval == 3 || interval == 4
}
