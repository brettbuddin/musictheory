# mt (music theory)

[![Build Status](https://travis-ci.org/brettbuddin/mt.svg?branch=master)](https://travis-ci.org/brettbuddin/mt)
[![GoDoc](https://godoc.org/github.com/brettbuddin/mt?status.svg)](https://godoc.org/github.com/brettbuddin/mt)

Explorations in music theory.

## Usage

```go
package main

import (
	"github.com/brettbuddin/mt"
)

func main() {
	root := mt.NewPitch(mt.C, 0, mt.Natural)

	root.Name(mt.SharpNames) // C4
	root.Freq()              // 261.625565 (Hz)
	root.MIDI()              // 72

	P5 := mt.Perfect(5)   // Perfect 5th
	A4 := mt.Augmented(4) // Augmented 4th

	root.Transpose(P5).Name(mt.SharpNames)          // G4
	root.Transpose(A4).Name(mt.SharpNames)          // F#4
	root.Transpose(P5.Negate()).Name(mt.SharpNames) // F3

	mt.NewScale(root, mt.DorianIntervals)
	// [C4, D4, Eb4, F4, G4, A4, Bb4]

	mt.NewScale(root, mt.MixolydianIntervals)
	// [C4, D4, E4, F4, G4, A4, Bb4]

    note := mt.NewNote(mt.C, 0, mt.Sharp, mt.D16) // C#4 sixteenth note
    note.Time(mt.D4, 120)                         // 125ms (quarter note getting the beat at 120 BPM)
}
```

楽しみます！

## TODO

- Chords
- Inversions
- Scientific pitch notation parsing
- ...
- Write a tool or two that uses this stuff...
