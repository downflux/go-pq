# go-pq
Golang generic priority queue implementation

This priority queue may be backed by either a min or max heap, depending on user specification.

This queue is not thread-safe. The user must add their own lock wrappers around this queue.

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

	qmin.Push(
		/* data = */ 418,
		/* priority = */ 10,
	)
	qmax.Push(418, 10)
	qmin.Push(42, 1)
	qmax.Push(42, 1)

	if data, p := qmin.Pop(); p != 1 || data != 42 {
		panic(fmt.Sprintf("found an unexpected value from pq.Pop(): %v, %v", p, data))
	}
	if data, p := qmax.Pop(); p != 10 || data != 418 {
		panic(fmt.Sprintf("found an unexpected value from pq.Pop(): %v, %v", p, data))
	}
}
```
