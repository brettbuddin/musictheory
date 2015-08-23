package math

import "math"

func Mod(n, m float64) float64 {
	out := math.Mod(n, m)
	if out < 0 {
		out += m
	}
	return out
}
