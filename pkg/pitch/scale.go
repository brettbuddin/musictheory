package pitch

import (
	"github.com/brettbuddin/mt/pkg/interval"
)

// Scales
var (
	ChromaticScale,
	MajorScale,
	MinorScale,
	IonianScale,
	DorianScale,
	PhrygianScale,
	AeolianScale,
	LydianScale,
	MixolydianScale,
	LocrianScale []interval.Interval
)

func init() {
	P1 := interval.Perfect(1)
	P4 := interval.Perfect(4)
	P5 := interval.Perfect(5)

	M2 := interval.Major(2)
	M3 := interval.Major(3)
	M6 := interval.Major(6)
	M7 := interval.Major(7)

	m2 := interval.Minor(2)
	m3 := interval.Minor(3)
	m6 := interval.Minor(6)
	m7 := interval.Minor(7)

	A4 := interval.Augmented(4)
	d5 := interval.Diminished(5)

	ChromaticScale = []interval.Interval{P1, m2, M2, m3, M3, P4, A4, P5, m6, M6, m7, M7}

	IonianScale = []interval.Interval{P1, M2, M3, P4, P5, M6, M7}
	MajorScale = IonianScale

	DorianScale = []interval.Interval{P1, M2, m3, P4, P5, m6, m7}
	PhrygianScale = []interval.Interval{P1, m2, m3, P4, P5, m6, m7}
	LydianScale = []interval.Interval{P1, M2, M3, A4, P5, M6, M7}
	MixolydianScale = []interval.Interval{P1, M2, M3, P4, P5, M6, m7}

	AeolianScale = []interval.Interval{P1, M2, m3, P4, P5, m6, m7}
	MinorScale = AeolianScale

	LocrianScale = []interval.Interval{P1, m2, m3, P4, d5, m6, m7}
}
