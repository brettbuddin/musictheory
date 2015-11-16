package mt

import (
	"fmt"
	"regexp"
	"strconv"
)

var notation = regexp.MustCompile("([ABCDEFG])(bb|b|#|x)?(\\d+)")

// MustParsePitch parses and returns a Pitch in scientific pitch notation or panics
func MustParsePitch(str string) *Pitch {
	pitch, err := ParsePitch(str)
	if err != nil {
		panic(err)
	}
	return pitch
}

// ParsePitch parses and returns a Pitch in scientific pitch notation
func ParsePitch(str string) (*Pitch, error) {
	matches := notation.FindStringSubmatch(str)
	if len(matches) < 1 {
		return nil, fmt.Errorf("no matches found")
	}

	class := matches[1]
	modifier := matches[2]
	octave, _ := strconv.Atoi(matches[3])

	classIndex, err := classNameIndex(class)
	if err != nil {
		return nil, err
	}

	modifierOffset, err := modifierNameOffset(modifier)
	if err != nil {
		return nil, err
	}

	pitch := NewPitch(classIndex+1, modifierOffset, octave)

	return &pitch, nil
}

func classNameIndex(name string) (int, error) {
	for i, n := range pitchNames {
		if n == name {
			return i, nil
		}
	}

	return 0, fmt.Errorf("unknown class name: %s", name)
}

func modifierNameOffset(name string) (int, error) {
	for i, a := range modifierNames {
		if a == name {
			return i - 2, nil
		}
	}

	return 0, fmt.Errorf("unknown modifier: %s", name)
}
