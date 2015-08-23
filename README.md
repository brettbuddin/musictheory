# mt (music theory)

[![Build Status](https://travis-ci.org/brettbuddin/mt.svg?branch=master)](https://travis-ci.org/brettbuddin/mt)

Explorations in music theory. Consider this a toy.

## Example

```go
package main

import (
	"fmt"
	mi "github.com/brettbuddin/mt/pkg/interval"
	mp "github.com/brettbuddin/mt/pkg/pitch"
)

var nameFmt = mp.SharpNames

func main() {
	tonic := mp.New(mp.C, 0, mp.Natural)
	fifthBelow := tonic.AddInterval(mi.Perfect(5).Negate())
	fifthAbove := tonic.AddInterval(mi.Perfect(5))

	format := "%s is a perfect fifth %s %s, and is %f Hz (MIDI tone %d).\n"

	fmt.Printf(format,
		fifthBelow.Name(nameFmt),
		"below",
		tonic.Name(nameFmt),
		fifthBelow.Freq(),
		fifthBelow.MIDI())

	fmt.Printf(format,
		fifthAbove.Name(nameFmt),
		"above",
		tonic.Name(nameFmt),
		fifthAbove.Freq(),
		fifthAbove.MIDI())
}
```

Outputs:
```
F3 is a perfect fifth below C4, and is 174.614116 Hz (MIDI tone 53).
G4 is a perfect fifth above C4, and is 391.995436 Hz (MIDI tone 67).
```

## TODO

- Document all of this
- Full test coverage
- Pitch Vectors
- Chords
- Scales (modal and special)
- Inversions
- Scientific pitch notation parsing
- ...
- Write a tool or two that uses this stuff...
