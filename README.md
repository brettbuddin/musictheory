# mt (music theory)

[![Build Status](https://travis-ci.org/brettbuddin/mt.svg?branch=master)](https://travis-ci.org/brettbuddin/mt)
[![GoDoc](https://godoc.org/github.com/brettbuddin/mt?status.svg)](https://godoc.org/github.com/brettbuddin/mt)

Explorations in music theory. Consider this a toy.

## Example

```go
package main

import (
	"fmt"
	mi "github.com/brettbuddin/mt/pkg/interval"
	mp "github.com/brettbuddin/mt/pkg/pitch"
)

func main() {
	origin := mp.New(mp.C, 0, mp.Natural)

	fmt.Printf("Origin: %s (%f Hz / MIDI %d)\n", 
            origin.Name(mp.SharpNames), 
            origin.Freq(), 
            origin.MIDI())

	fmt.Println("Perfect fifth below:", origin.Transpose(mi.Perfect(5).Negate()).Name(mp.SharpNames))
	fmt.Println("Perfect fifth above:", origin.Transpose(mi.Perfect(5)).Name(mp.SharpNames))
	fmt.Println("Augmented fourth above:", origin.Transpose(mi.Augmented(4)).Name(mp.SharpNames))
}
```

Outputs:
```
Origin: C4 (261.625565 Hz / MIDI 60)
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
