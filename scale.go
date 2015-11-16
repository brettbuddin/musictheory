package mt

// Scales
var (
	ChromaticIntervals  []Interval
	MajorIntervals      []Interval
	MinorIntervals      []Interval
	IonianIntervals     []Interval
	DorianIntervals     []Interval
	PhrygianIntervals   []Interval
	AeolianIntervals    []Interval
	LydianIntervals     []Interval
	MixolydianIntervals []Interval
	LocrianIntervals    []Interval
)

func init() {
	P1 := Perfect(1)
	P4 := Perfect(4)
	P5 := Perfect(5)

	M2 := Major(2)
	M3 := Major(3)
	M6 := Major(6)
	M7 := Major(7)

	m2 := Minor(2)
	m3 := Minor(3)
	m6 := Minor(6)
	m7 := Minor(7)

	A4 := Augmented(4)
	d5 := Diminished(5)

	ChromaticIntervals = []Interval{P1, m2, M2, m3, M3, P4, A4, P5, m6, M6, m7, M7}

	IonianIntervals = []Interval{P1, M2, M3, P4, P5, M6, M7}
	MajorIntervals = IonianIntervals

	DorianIntervals = []Interval{P1, M2, m3, P4, P5, M6, m7}
	PhrygianIntervals = []Interval{P1, m2, m3, P4, P5, m6, m7}
	LydianIntervals = []Interval{P1, M2, M3, A4, P5, M6, M7}
	MixolydianIntervals = []Interval{P1, M2, M3, P4, P5, M6, m7}

	AeolianIntervals = []Interval{P1, M2, m3, P4, P5, m6, m7}
	MinorIntervals = AeolianIntervals

	LocrianIntervals = []Interval{P1, m2, m3, P4, d5, m6, m7}
}

// Scale is a series of Transposers
type Scale []Transposer

// Transpose transposes a scale by the specified Interval
func (s Scale) Transpose(i Interval) Transposer {
	scale := Scale{}
	for _, transposer := range s {
		scale = append(scale, transposer.Transpose(i))
	}
	return scale
}

// NewScale returns a Scale built using a set of intervals
func NewScale(root Transposer, intervals []Interval) Scale {
	scale := Scale{}
	for _, i := range intervals {
		scale = append(scale, root.Transpose(i))
	}
	return scale
}
