package musictheory

func normalizeChromatic(v int) int {
	return v % 12
}

func normalizeDiatonic(v int) int {
	return v % 7
}

func normalizeChromaticPositive(v int) int {
	r := v % 12
	if r < 0 {
		r += 12
	}
	return r
}

func diatonicOctaves(v int) int {
	return v / 7
}

func chromaticOctaves(v int) int {
	return v / 12
}

func inverseChromatic(v int) int {
	return 12 - v
}

func inverseDiatonic(v int) int {
	return 7 - v
}
