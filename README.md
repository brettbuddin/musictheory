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

func main() {
	tonic := mp.New(mp.C, 0, mp.Natural)
	fifthBelow := tonic.AddInterval(mi.Perfect(5).Negate())
	fifthAbove := tonic.AddInterval(mi.Perfect(5))
	names := mp.SharpNames

	fmt.Printf("%s is a perfect fifth below %s\n", fifthBelow.Name(names), tonic.Name(names))
	fmt.Printf("%s is a perfect fifth above %s\n", fifthAbove.Name(names), tonic.Name(names))
}
```

Outputs:
```
G0 is a perfect fifth above C0
F-1 is a perfect fifth above C0
```
