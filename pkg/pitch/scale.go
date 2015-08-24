package pitch

import (
	"github.com/brettbuddin/mt/pkg/interval"
)

// PitchVector is a series of pitches
type PitchSeries []Pitch

// Scale returns a series of pitches separated by the provided scale's intervals
func Scale(root Pitch, scale interval.Scale) PitchSeries {
	series := PitchSeries{}
	for _, i := range scale {
		series = append(series, root.Transpose(i))
	}
	return series
}
