package pq

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type X struct {
	X int
}

var _ heap.Interface = &max[X]{}

func TestPQ(t *testing.T) {
	type item struct {
		p        X
		priority float64
	}

	type config struct {
		name string
		data []item
		size int
		want []X
	}

	configs := []config{
		{
			name: "Null",
			data: nil,
			want: nil,
		},

		{
			name: "Trivial",
			data: []item{
				{
					p:        X{1},
					priority: 0,
				},
			},
			size: 1,
			want: []X{
				X{1},
			},
		},
		{
			name: "Trivial/NoSizeLimit",
			data: []item{
				{
					p:        X{1},
					priority: 0,
				},
			},
			size: 0,
			want: []X{
				X{1},
			},
		},
		{
			name: "Sorted",
			data: []item{
				{
					p:        X{0},
					priority: 5,
				},
				{
					p:        X{-1},
					priority: 4,
				},
				{
					p:        X{-2},
					priority: 3,
				},
				{
					p:        X{-3},
					priority: 2,
				},
				{
					p:        X{-4},
					priority: 1,
				},
			},
			size: 5,
			want: []X{
				X{0},
				X{-1},
				X{-2},
				X{-3},
				X{-4},
			},
		},
		{
			name: "Sorted/Reverse",
			data: []item{
				{
					p:        X{0},
					priority: 1,
				},
				{
					p:        X{-1},
					priority: 2,
				},
				{
					p:        X{-2},
					priority: 3,
				},
				{
					p:        X{-3},
					priority: 4,
				},
				{
					p:        X{-4},
					priority: 5,
				},
			},
			size: 5,
			want: []X{
				X{-4},
				X{-3},
				X{-2},
				X{-1},
				X{0},
			},
		},
		{
			name: "Sorted/Shuffled",
			data: []item{
				{
					p:        X{0},
					priority: 4,
				},
				{
					p:        X{-1},
					priority: 2,
				},
				{
					p:        X{-2},
					priority: 5,
				},
				{
					p:        X{-3},
					priority: 1,
				},
				{
					p:        X{-4},
					priority: 3,
				},
			},
			size: 5,
			want: []X{
				X{-2},
				X{0},
				X{-4},
				X{-1},
				X{-3},
			},
		},
	}

	for _, c := range configs {
		t.Run(fmt.Sprintf("Max/%v", c.name), func(t *testing.T) {
			pq := New[X](c.size, PMax)
			for _, d := range c.data {
				pq.Push(d.p, d.priority)
			}

			var got []X
			for !pq.Empty() {
				got = append(got, pq.Pop())
			}

			if diff := cmp.Diff(
				c.want, got,
			); diff != "" {
				t.Errorf("Pop() mismatch (-want +got):\n%v", diff)
			}
		})
		// Run automated tests using a min-PQ instead. This test is
		// auto-generated.
		t.Run(fmt.Sprintf("Min/%v", c.name), func(t *testing.T) {
			var want []X
			for i := 0; i < len(c.want); i++ {
				want = append(want, c.want[len(c.want)-i-1])
			}

			pq := New[X](c.size, PMin)
			for _, d := range c.data {
				pq.Push(d.p, d.priority)
			}

			var got []X
			for !pq.Empty() {
				got = append(got, pq.Pop())
			}

			if diff := cmp.Diff(
				want, got,
			); diff != "" {
				t.Errorf("Pop() mismatch (-want +got):\n%v", diff)
			}
		})
	}
}
