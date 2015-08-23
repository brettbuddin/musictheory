# mt (music theory)

[![Build Status](https://travis-ci.org/brettbuddin/mt.svg?branch=master)](https://travis-ci.org/brettbuddin/mt)

Explorations in music theory. Consider this a toy.

### Example

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

	format := "%s is a perfect fifth %s %s, and is %f Hz.\n"
	fmt.Printf(format, fifthBelow.Name(nameFmt), "below", tonic.Name(nameFmt), fifthBelow.Freq())
	fmt.Printf(format, fifthAbove.Name(nameFmt), "above", tonic.Name(nameFmt), fifthAbove.Freq())
}
```

Outputs:
```
F3 is a perfect fifth below C4, and is 174.614116 Hz.
G4 is a perfect fifth above C4, and is 391.995436 Hz.
```

