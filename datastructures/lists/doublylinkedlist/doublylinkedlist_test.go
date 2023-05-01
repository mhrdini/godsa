package doublylinkedlist

import (
	"fmt"
	"testing"

	"github.com/mhrdini/godsa/helpers"
)

func TestInsertAt(t *testing.T) {

	list := []int{1, 2, 3, 4, 5}
	inserted := []int{0, 0, 0}

	testCases := []struct {
		index int
		want  string
	}{
		{0, helpers.ToString([]int{0, 0, 0, 1, 2, 3, 4, 5})},
		{1, helpers.ToString([]int{1, 0, 0, 0, 2, 3, 4, 5})},
		{4, helpers.ToString([]int{1, 2, 3, 4, 0, 0, 0, 5})},
		{5, helpers.ToString([]int{1, 2, 3, 4, 5, 0, 0, 0})},
	}

	for _, tc := range testCases {
		list := New(list...)
		list.InsertAt(tc.index, inserted...)
		got := list.String()
		want := tc.want
		helpers.AssertEqual(t, got, want)
	}
}

func TestRemove(t *testing.T) {

	empty := []int{}
	arbitrary := []int{1, 2, 3}

	t.Run("any index on empty list", func(t *testing.T) {
		empty := New(empty...)
		_, ok := empty.Remove(0)
		got := ok
		want := false
		helpers.AssertEqual(t, got, want)
	})

	type result struct {
		value int
		ok    bool
		list  string
	}

	testCases := []struct {
		index int
		want  result
	}{
		{0, result{arbitrary[0], true, helpers.ToString([]int{2, 3})}},
		{1, result{arbitrary[1], true, helpers.ToString([]int{1, 3})}},
		{2, result{arbitrary[2], true, helpers.ToString([]int{1, 2})}},
		{3, result{0, false, helpers.ToString([]int{1, 2, 3})}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("index %d on arbitrary list %v", tc.index, arbitrary), func(t *testing.T) {
			arbitrary := New(arbitrary...)
			value, ok := arbitrary.Remove(tc.index)
			helpers.AssertEqual(t, value, tc.want.value)
			helpers.AssertEqual(t, ok, tc.want.ok)
			helpers.AssertEqual(t, arbitrary.String(), tc.want.list)
		})
	}
}
