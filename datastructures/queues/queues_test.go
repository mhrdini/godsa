package queues

import (
	"fmt"
	"testing"

	"github.com/mhrdini/godsa/datastructures/queues/linkedlistqueue"
	"github.com/mhrdini/godsa/helpers"
)

func queues[T any](base ...T) []Queue[T] {
	return []Queue[T]{
		linkedlistqueue.New(base...),
	}
}

func TestEnqueue(t *testing.T) {

	empty := []int{}
	arbitrary := []int{1, 2, 3}

	testCases := []struct {
		base  []int
		input []int
		want  string
	}{
		{empty, empty, "[]"},
		{empty, arbitrary, "[1 2 3]"},
		{arbitrary, arbitrary, "[1 2 3 1 2 3]"},
	}

	for _, tc := range testCases {
		for _, queue := range queues(tc.base...) {
			t.Run(fmt.Sprintf("to %v %v <- %v", queue.Name(), tc.base, tc.input), func(t *testing.T) {
				for _, v := range tc.input {
					queue.Enqueue(v)
				}
				helpers.AssertEqual(t, queue.String(), tc.want)
			})
		}
	}
}

func TestDequeue(t *testing.T) {
	empty := []int{}
	arbitrary := []int{1, 2, 3}

	type result struct {
		value  int
		ok     bool
		values string
	}

	testCases := []struct {
		base []int
		want result
	}{
		{empty, result{0, false, "[]"}},
		{arbitrary, result{1, true, "[2 3]"}},
	}

	for _, tc := range testCases {
		for _, queue := range queues(tc.base...) {
			t.Run(fmt.Sprintf("on %v %v", queue.Name(), tc.base), func(t *testing.T) {
				value, ok := queue.Dequeue()
				helpers.AssertEqual(t, tc.want.value, value)
				helpers.AssertEqual(t, tc.want.ok, ok)
				helpers.AssertEqual(t, tc.want.values, queue.String())
			})
		}
	}
}

func TestPeek(t *testing.T) {
	empty := []int{}
	arbitrary := []int{1, 2, 3}

	type result struct {
		value  int
		ok     bool
		values string
	}

	testCases := []struct {
		base []int
		want result
	}{
		{empty, result{0, false, "[]"}},
		{arbitrary, result{1, true, "[1 2 3]"}},
	}

	for _, tc := range testCases {
		for _, queue := range queues(tc.base...) {
			t.Run(fmt.Sprintf("on %v %v", queue.Name(), tc.base), func(t *testing.T) {
				value, ok := queue.Peek()
				helpers.AssertEqual(t, tc.want.value, value)
				helpers.AssertEqual(t, tc.want.ok, ok)
				helpers.AssertEqual(t, tc.want.values, queue.String())
			})
		}
	}
}
