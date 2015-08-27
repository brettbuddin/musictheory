package mt

// Scales
var (
	ChromaticIntervals  Intervals
	MajorIntervals      Intervals
	MinorIntervals      Intervals
	IonianIntervals     Intervals
	DorianIntervals     Intervals
	PhrygianIntervals   Intervals
	AeolianIntervals    Intervals
	LydianIntervals     Intervals
	MixolydianIntervals Intervals
	LocrianIntervals    Intervals
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

	ChromaticIntervals = Intervals{P1, m2, M2, m3, M3, P4, A4, P5, m6, M6, m7, M7}

	IonianIntervals = Intervals{P1, M2, M3, P4, P5, M6, M7}
	MajorIntervals = IonianIntervals

	DorianIntervals = Intervals{P1, M2, m3, P4, P5, M6, m7}
	PhrygianIntervals = Intervals{P1, m2, m3, P4, P5, m6, m7}
	LydianIntervals = Intervals{P1, M2, M3, A4, P5, M6, M7}
	MixolydianIntervals = Intervals{P1, M2, M3, P4, P5, M6, m7}

	AeolianIntervals = Intervals{P1, M2, m3, P4, P5, m6, m7}
	MinorIntervals = AeolianIntervals

	LocrianIntervals = Intervals{P1, m2, m3, P4, d5, m6, m7}
}

// Scale is a series of pitches
type Scale []Pitch

// Transpose transposes a scale by the specified Interval
func (s Scale) Transpose(i Interval) Scale {
	scale := []Pitch{}
	for _, pitch := range s {
		scale = append(scale, pitch.Transpose(i))
	}
	return scale
}

// NewScale returns a Scale built using a set of intervals
func NewScale(root Pitch, intervals Intervals) Scale {
	scale := Scale{}
	for _, i := range intervals {
		scale = append(scale, root.Transpose(i))
	}
	return scale
}
