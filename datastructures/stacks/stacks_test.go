package stacks

import (
	"fmt"
	"testing"

	"github.com/mhrdini/godsa/datastructures/stacks/arraystack"
	"github.com/mhrdini/godsa/datastructures/stacks/linkedliststack"
	"github.com/mhrdini/godsa/helpers"
)

func stacks[T any](base ...T) []Stack[T] {
	return []Stack[T]{
		arraystack.New(base...),
		linkedliststack.New(base...),
	}
}

func TestPush(t *testing.T) {

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
		for _, stack := range stacks(tc.base...) {
			t.Run(fmt.Sprintf("to %v %v <- %v", stack.Name(), tc.base, tc.input), func(t *testing.T) {
				for _, v := range tc.input {
					stack.Push(v)
				}
				helpers.AssertEqual(t, stack.String(), tc.want)
			})
		}
	}
}

func TestPop(t *testing.T) {
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
		{arbitrary, result{3, true, "[1 2]"}},
	}

	for _, tc := range testCases {
		for _, stack := range stacks(tc.base...) {
			t.Run(fmt.Sprintf("on %v %v", stack.Name(), tc.base), func(t *testing.T) {
				value, ok := stack.Pop()
				helpers.AssertEqual(t, tc.want.value, value)
				helpers.AssertEqual(t, tc.want.ok, ok)
				helpers.AssertEqual(t, tc.want.values, stack.String())
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
		{arbitrary, result{3, true, "[1 2 3]"}},
	}

	for _, tc := range testCases {
		for _, stack := range stacks(tc.base...) {
			t.Run(fmt.Sprintf("on %v %v", stack.Name(), tc.base), func(t *testing.T) {
				value, ok := stack.Peek()
				helpers.AssertEqual(t, tc.want.value, value)
				helpers.AssertEqual(t, tc.want.ok, ok)
				helpers.AssertEqual(t, tc.want.values, stack.String())
			})
		}
	}
}
