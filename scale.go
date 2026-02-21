package musictheory

import "slices"

// Scale is a series of Pitches
type Scale []Pitch

// Transpose transposes a scale by the specified Interval
func (s Scale) Transpose(i Interval) Scale {
	scale := make(Scale, len(s))
	for j, p := range s {
		scale[j] = p.Transpose(i)
	}
	return scale
}

// NewScale returns a Scale built using a set of intervals
func NewScale(root Pitch, intervals []Interval, octaves int) Scale {
	descending := octaves < 0
	n := octaves
	if n < 0 {
		n = -n
	}

	scale := make(Scale, 0, n*len(intervals))
	originalRoot := root

	// Begin at the base of our octave shift
	if descending {
		root = root.Transpose(Octave(octaves))
	}

	for i := 0; i < n; i++ {
		for j, v := range intervals {
			// Ignore the tonic which will become the *last* item in the slice
			// once reversed. This is to maintain consistency with ascending
			// scales: they don't include the final octave of the tonic.
			if descending && i == 0 && j == 0 {
				continue
			}
			scale = append(scale, root.Transpose(v))
		}
		root = root.Transpose(Octave(1))
	}

	// Add the original tonic to the end. It's about to become the beginning of
	// the slice once it's reversed. Reversing the list produces our descending
	// scale.
	if descending {
		scale = append(scale, originalRoot)
		slices.Reverse(scale)
	}

	return scale
}
