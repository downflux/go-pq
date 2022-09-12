# go-pq
Golang generic priority queue implementation

This priority queue may be backed by either a min or max heap, depending on user specification.

## Example

```golang
package main

import (
	"fmt"

	"github.com/downflux/go-pq/pq"
)

type F float64

func main() {
	qmin := pq.New[F](2, pq.PMin)
	qmax := pq.New[F](2, pq.PMax)

	qmin.Push( /* data = */ 418 /* priority = */, 10)
	qmax.Push(418, 10)
	qmin.Push(42, 1)
	qmax.Push(42, 1)

	if p := qmin.Priority(); p != 1 {
		panic(fmt.Sprintf("found an unexpected priority for the next element in the priority queue: %v", p))
	}
	if p := qmax.Priority(); p != 10 {
		panic(fmt.Sprintf("found an unexpected priority for the next element in the priority queue: %v", p))
	}
	if data := qmin.Pop(); data != 42 {
		panic(fmt.Sprintf("did not find the correct value from the queue: %v", data))
	}
	if data := qmax.Pop(); data != 418 {
		panic(fmt.Sprintf("did not find the correct value from the queue: %v", data))
	}
}
```
