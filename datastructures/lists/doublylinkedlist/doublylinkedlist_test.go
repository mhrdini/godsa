package doublylinkedlist

import (
	"fmt"
	"testing"

	"github.com/mhrdini/godsa/helpers"
)

func TestNew(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		empty := []int{}
		got := New(empty...).Empty()
		want := true
		helpers.AssertEqual(t, got, want)
	})

	t.Run("arbitrary list", func(t *testing.T) {
		list := []int{1, 2, 3}
		got := New(list...).String()
		want := fmt.Sprint(list)
		helpers.AssertEqual(t, got, want)
	})
}

func TestConcat(t *testing.T) {

	emptyBase := []int{}
	arbitraryBase := []int{1, 2, 3}

	emptyList := []int{}
	arbitraryList := []int{4, 5, 6}

	testCases := []struct {
		input       [][]int
		onEmpty     string
		onArbitrary string
	}{
		{[][]int{}, helpers.ToString(emptyBase), helpers.ToString(arbitraryBase)},
		{[][]int{emptyList, arbitraryList}, helpers.ToString(arbitraryList), helpers.ToString(append(arbitraryBase, arbitraryList...))},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v on empty list %v", tc.input, emptyBase), func(t *testing.T) {
			list := New(emptyBase...)
			input := helpers.Map(tc.input, func(l []int) *List[int] { return New(l...) })
			list.Concat(input...)
			helpers.AssertEqual(t, helpers.ToString(list), tc.onEmpty)
		})

		t.Run(fmt.Sprintf("%v on arbitrary list %v", tc.input, arbitraryBase), func(t *testing.T) {
			list := New(arbitraryBase...)
			input := helpers.Map(tc.input, func(l []int) *List[int] { return New(l...) })
			list.Concat(input...)
			helpers.AssertEqual(t, helpers.ToString(list), tc.onArbitrary)
		})
	}
}
