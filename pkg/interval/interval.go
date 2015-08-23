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

const (
	Unison int = iota + 1
	Second
	Third
	Fourth
	Fifth
	Sixth
	Seventh
	Octave
	Ninth
	Tenth
	Eleventh
	Twelfth
	Thirteenth
	Fourteenth
	Fiftheenth
)

var (
	Perfect          = buildFromQuality(Quality{PerfectT, 0})
	Major            = buildFromQuality(Quality{MajorT, 0})
	Minor            = buildFromQuality(Quality{MinorT, 0})
	Augmented        = buildFromQuality(Quality{AugmentedT, 1})
	DoublyAugmented  = buildFromQuality(Quality{AugmentedT, 2})
	Diminished       = buildFromQuality(Quality{DiminishedT, 1})
	DoublyDiminished = buildFromQuality(Quality{DiminishedT, 2})
)

func buildFromQuality(quality Quality) func(int) Interval {
	return func(val int) Interval {
		diatonic := int(mt_math.Mod(float64(val-1), 7))
		diff := qualityToDiff(perfect(diatonic), quality)
		return build(val, diff)
	}
}

func build(val, offset int) Interval {
	diatonic := int(mt_math.Mod(float64(val-1), 7))
	return Interval{
		octaves:   int((val - 1) / 7.0),
		diatonic:  diatonic,
		chromatic: diatonicToChromatic(diatonic) + offset,
	}
}

func New(val, offset int) Interval {
	return build(val, offset)
}

type Interval struct {
	octaves   int
	diatonic  int
	chromatic int
}

func (i Interval) Octaves() int {
	return i.octaves
}

func (i Interval) Diff() int {
	return i.chromatic - diatonicToChromatic(i.diatonic)
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

type Quality struct {
	Type, Size int
}

func diatonicToChromatic(interval int) int {
	if interval >= len(diatonicToChromaticLookup) {
		panic(fmt.Sprintf("interval out of range: %d", interval))
	}

	return diatonicToChromaticLookup[interval]
}

var diatonicToChromaticLookup = []int{0, 2, 4, 5, 7, 9, 11}

func qualityToDiff(perfect bool, q Quality) int {
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

func diffToQuality(perfect bool, diff int) Quality {
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
