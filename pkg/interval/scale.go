package interval

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
	LocrianScale []Interval
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

	ChromaticScale = []Interval{P1, m2, M2, m3, M3, P4, A4, P5, m6, M6, m7, M7}

	IonianScale = []Interval{P1, M2, M3, P4, P5, M6, M7}
	MajorScale = IonianScale

	DorianScale = []Interval{P1, M2, m3, P4, P5, m6, m7}
	PhrygianScale = []Interval{P1, m2, m3, P4, P5, m6, m7}
	LydianScale = []Interval{P1, M2, M3, A4, P5, M6, M7}
	MixolydianScale = []Interval{P1, M2, M3, P4, P5, M6, m7}

	AeolianScale = []Interval{P1, M2, m3, P4, P5, m6, m7}
	MinorScale = AeolianScale

	LocrianScale = []Interval{P1, m2, m3, P4, d5, m6, m7}
}
