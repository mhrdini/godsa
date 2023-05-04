package deque

import (
	"fmt"
	"testing"

	"github.com/mhrdini/godsa/helpers"
)

func emptyList() []int {
	return []int{}
}

func arbitraryList() []int {
	return []int{1, 2, 3}
}

type result struct {
	value  int
	ok     bool
	values string
}

func TestPush(t *testing.T) {
	testCases := []struct {
		base  []int
		input []int
		want  string
	}{
		{emptyList(), emptyList(), "[]"},
		{emptyList(), arbitraryList(), "[3 2 1]"},
		{arbitraryList(), emptyList(), "[1 2 3]"},
		{arbitraryList(), arbitraryList(), "[3 2 1 1 2 3]"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v -> %v", tc.input, tc.base), func(t *testing.T) {
			deque := New(tc.base...)
			for _, v := range tc.input {
				deque.Push(v)
			}
			helpers.AssertEqual(t, deque.String(), tc.want)
		})
	}
}

func Insert(t *testing.T) {
	testCases := []struct {
		base  []int
		input []int
		want  string
	}{
		{emptyList(), emptyList(), "[]"},
		{emptyList(), arbitraryList(), "[1 2 3]"},
		{arbitraryList(), emptyList(), "[1 2 3]"},
		{arbitraryList(), arbitraryList(), "[1 2 3 1 2 3]"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v <- %v", tc.base, tc.input), func(t *testing.T) {
			deque := New(tc.base...)
			for _, v := range tc.input {
				deque.Push(v)
			}
			helpers.AssertEqual(t, deque.String(), tc.want)
		})
	}
}

func TestPop(t *testing.T) {
	testCases := []struct {
		base []int
		want result
	}{
		{emptyList(), result{0, false, "[]"}},
		{arbitraryList(), result{1, true, "[2 3]"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("<- %v", tc.base), func(t *testing.T) {
			deque := New(tc.base...)
			v, ok := deque.Pop()
			helpers.AssertEqual(t, v, tc.want.value)
			helpers.AssertEqual(t, ok, tc.want.ok)
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		base []int
		want result
	}{
		{emptyList(), result{0, false, "[]"}},
		{arbitraryList(), result{3, true, "[1 2]"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v ->", tc.base), func(t *testing.T) {
			deque := New(tc.base...)
			v, ok := deque.Remove()
			helpers.AssertEqual(t, v, tc.want.value)
			helpers.AssertEqual(t, ok, tc.want.ok)
		})
	}
}
