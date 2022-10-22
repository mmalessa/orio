Package orio provides GPIO access on the OrangePi Zero  
(c) Marcin Malessa https://github.com/mmalessa/orio

Early version - use at your own risk.

```go
package main

import (
	"time"

	"github.com/mmalessa/orio"
)

func main() {
	orio.DebugMode = true
	defer orio.Close()

	led1 := orio.Pin(orio.PH7)

	checkError(led1.Output())

	checkError(led1.High())
	time.Sleep(500 * time.Millisecond)
	checkError(led1.Low())
	time.Sleep(500 * time.Millisecond)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
```

