# mt (music theory)

[![Build Status](https://travis-ci.org/brettbuddin/mt.svg?branch=master)](https://travis-ci.org/brettbuddin/mt)
[![GoDoc](https://godoc.org/github.com/brettbuddin/mt?status.svg)](https://godoc.org/github.com/brettbuddin/mt)

Explorations in music theory. Consider this a toy.

## Example

```go
package main

import (
	"fmt"
	"github.com/brettbuddin/mt"
)

func main() {
	origin := mt.NewPitch(mt.C, 0, mt.Natural)

	fmt.Printf("Origin: %s (%f Hz / MIDI %d)\n", 
            origin.Name(mt.SharpNames), 
            origin.Freq(), 
            origin.MIDI())

    P5 := mt.Perfect(5)
    A4 := mt.Augmented(4)

	fmt.Println("Perfect fifth below:", origin.Transpose(P5.Negate()).Name(mt.SharpNames))
	fmt.Println("Perfect fifth above:", origin.Transpose(P5).Name(mt.SharpNames))
	fmt.Println("Augmented fourth above:", origin.Transpose(A4).Name(mt.SharpNames))
}
```

Outputs:
```
Origin: C4 (261.625565 Hz / MIDI 72)
Perfect fifth below: F3
Perfect fifth above: G4
Augmented fourth above: F#4
```

## TODO

- Squash bugs
- Document all of this
- Full test coverage
- Pitch Vectors
- Chords
- Scales (modal and special)
- Inversions
- Scientific pitch notation parsing
- ...
- Write a tool or two that uses this stuff...
