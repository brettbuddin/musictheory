# mt (music theory)

[![Build Status](https://travis-ci.org/brettbuddin/mt.svg?branch=master)](https://travis-ci.org/brettbuddin/mt)

Explorations in music theory. Consider this a toy.

```go
package main

import (
	"fmt"
	mi "github.com/brettbuddin/mt/pkg/interval"
	mp "github.com/brettbuddin/mt/pkg/pitch"
)

func main() {
	tonic := mp.New(mp.C, 0, mp.Natural)
	perfectFifth := tonic.AddInterval(mi.Perfect(5))
	names := mp.SharpNames

	fmt.Printf("%s is a perfect fifth above %s\n", perfectFifth.Name(names), tonic.Name(names))
}
```

Outputs:
```
// G is a perfect fifth above C
```
