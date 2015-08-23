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

	fmt.Printf("%s is a perfect fifth below %s\n", fifthBelow.Name(nameFmt), tonic.Name(nameFmt))
	fmt.Printf("%s is a perfect fifth above %s\n", fifthAbove.Name(nameFmt), tonic.Name(nameFmt))
}
```

Outputs:
```
G0 is a perfect fifth above C0
F-1 is a perfect fifth above C0
```
