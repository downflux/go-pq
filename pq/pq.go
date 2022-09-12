package pq

import (
	"container/heap"
	"math"
)

// P is the queue priority type, i.e. a queue backed by a min- or max-heap.
type P bool

const (
	PMin = false
	PMax = true
)

func (p P) P(w float64) float64 {
	return map[P]float64{
		PMin: -w,
		PMax: w,
	}[p]
}

type node[T any] struct {
	p      T
	index  int
	weight float64
}

// max implements the heap.Interface.
//
// See https://pkg.go.dev/container/heap for the heap implementation from which
// this was shamelessly stolen.
type max[T any] []*node[T]

func (heap *max[T]) Len() int           { return len(*heap) }
func (heap *max[T]) Less(i, j int) bool { return (*heap)[i].weight > (*heap)[j].weight }
func (heap *max[T]) Swap(i, j int) {
	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
	(*heap)[i].index = i
	(*heap)[j].index = j
}
func (heap *max[T]) Push(x any) {
	m := x.(*node[T])
	m.index = heap.Len()
	*heap = append(*heap, m)
}
func (heap *max[T]) Pop() any {
	h := *heap
	n := len(h)
	m := h[n-1]

	h[n-1] = nil
	m.index = -1

	*heap = h[0 : n-1]
	return m
}

// PQ is a priority queue of arbitrary data with a maximum queue size. Data
// popped from the queue is of highest priority. The priority queue may be a
// max or min queue.
type PQ[T any] struct {
	heap *max[T]
	n    int
	p    P
}

// New creates a priority queue instance. Size indicates the maximum size of the
// heap -- a size of 0 indicates no limit on the queue size.
func New[T any](size int, priority P) *PQ[T] {
	h := max[T](make([]*node[T], 0, int(math.Max(100.0, float64(size)))))
	pq := &PQ[T]{
		heap: &h,
		n:    size,
		p:    priority,
	}
	heap.Init(pq.heap)
	return pq
}

func (pq *PQ[T]) Len() int    { return pq.heap.Len() }
func (pq *PQ[T]) Empty() bool { return pq.Len() == 0 }
func (pq *PQ[T]) Full() bool  { return pq.n > 0 && pq.Len() >= pq.n }

// Priority calculates the current highest priority of the queue.
func (pq *PQ[T]) Priority() float64 {
	if pq.Empty() {
		return math.Inf(0)
	}

	// See https://groups.google.com/g/golang-nuts/c/sy1p8SfyPoY.
	return pq.p.P((*pq.heap)[0].weight)
}

// Push adds a new point into the queue.
//
// The queue will enforce the struct size constraint by removing elements frmo
// itself until the constraint is satisfied.
func (pq *PQ[T]) Push(p T, priority float64) {
	heap.Push(pq.heap, &node[T]{
		p:      p,
		weight: pq.p.P(priority),
	})
	for !pq.Empty() && pq.Len() > pq.n && pq.n > 0 {
		pq.Pop()
	}
}

// Pop removes the node with the highest priority from the queue.
func (pq *PQ[T]) Pop() T { return heap.Pop(pq.heap).(*node[T]).p }
