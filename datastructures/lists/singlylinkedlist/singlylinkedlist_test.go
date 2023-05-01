package singlylinkedlist

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
		assertEqual(t, got, want)
	})

	t.Run("arbitrary list", func(t *testing.T) {
		list := []int{1, 2, 3}
		got := New(list...).String()
		want := fmt.Sprint(list)
		assertEqual(t, got, want)
	})
}

func TestAdd(t *testing.T) {

	empty := []int{}
	arbitrary := []int{1, 2, 3}

	testCases := []struct {
		values          []int
		wantOnEmpty     int
		wantOnArbitrary int
	}{
		{[]int{}, len(empty), len(arbitrary)},
		{[]int{4, 5, 6}, len(empty) + 3, len(arbitrary) + 3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v onto empty list %v", tc.values, empty), func(t *testing.T) {
			list := New(empty...)
			list.Add(tc.values...)
			assertEqual(t, list.Size(), tc.wantOnEmpty)
		})

		t.Run(fmt.Sprintf("%v onto arbitrary list %v", tc.values, arbitrary), func(t *testing.T) {
			list := New(arbitrary...)
			list.Add(tc.values...)
			assertEqual(t, list.Size(), tc.wantOnArbitrary)
		})
	}
}

func TestInsertAt(t *testing.T) {
	baseLists := []struct {
		list           []int
		startPosition  int
		middlePosition int
		endPosition    int
		errorPosition  int
	}{
		{[]int{}, 0, 0, 0, 1},
		{[]int{1, 2, 3}, 0, 1, 3, 4},
	}

	emptyBase := baseLists[0]
	arbitraryBase := baseLists[1]

	emptyTest := []int{}
	arbitraryTest := []int{4, 5, 6}

	type wants struct {
		onStartPosition  string
		onMiddlePosition string
		onEndPosition    string
		onErrorPosition  string
	}

	testCases := []struct {
		list        []int
		onEmpty     wants
		onArbitrary wants
	}{
		{
			emptyTest,
			wants{
				onStartPosition:  toString(emptyBase.list),
				onMiddlePosition: toString(emptyBase.list),
				onEndPosition:    toString(emptyBase.list),
				onErrorPosition:  toString(emptyBase.list),
			},
			wants{
				onStartPosition:  toString(arbitraryBase.list),
				onMiddlePosition: toString(arbitraryBase.list),
				onEndPosition:    toString(arbitraryBase.list),
				onErrorPosition:  toString(arbitraryBase.list),
			},
		},
		{
			arbitraryTest,
			wants{
				onStartPosition:  toString(arbitraryTest),
				onMiddlePosition: toString(arbitraryTest),
				onEndPosition:    toString(arbitraryTest),
				onErrorPosition:  toString(emptyBase.list),
			},
			wants{
				onStartPosition:  toString([]int{4, 5, 6, 1, 2, 3}),
				onMiddlePosition: toString([]int{1, 4, 5, 6, 2, 3}),
				onEndPosition:    toString([]int{1, 2, 3, 4, 5, 6}),
				onErrorPosition:  toString([]int{1, 2, 3}),
			},
		},
	}

	for _, tc := range testCases {
		// for empty base list
		t.Run(fmt.Sprintf("%v into start position (%d) of empty list %v", toString(tc.list), emptyBase.startPosition, toString(emptyBase.list)), func(t *testing.T) {
			base := New(emptyBase.list...)
			base.InsertAt(emptyBase.startPosition, tc.list...)
			got := base.String()
			assertEqual(t, got, tc.onEmpty.onStartPosition)
		})

		t.Run(fmt.Sprintf("%v into middle position (%d) of empty list %v", toString(tc.list), emptyBase.middlePosition, toString(emptyBase.list)), func(t *testing.T) {
			base := New(emptyBase.list...)
			base.InsertAt(emptyBase.middlePosition, tc.list...)
			got := base.String()
			assertEqual(t, got, tc.onEmpty.onMiddlePosition)
		})

		t.Run(fmt.Sprintf("%v into end position (%d) of empty list %v", toString(tc.list), emptyBase.endPosition, toString(emptyBase.list)), func(t *testing.T) {
			base := New(emptyBase.list...)
			base.InsertAt(emptyBase.endPosition, tc.list...)
			got := base.String()
			assertEqual(t, got, tc.onEmpty.onEndPosition)
		})

		t.Run(fmt.Sprintf("%v into error position (%d) of empty list %v", toString(tc.list), emptyBase.errorPosition, toString(emptyBase.list)), func(t *testing.T) {
			base := New(emptyBase.list...)
			base.InsertAt(emptyBase.errorPosition, tc.list...)
			got := base.String()
			assertEqual(t, got, tc.onEmpty.onErrorPosition)
		})

		// for arbitrary base list
		t.Run(fmt.Sprintf("%v into start position (%d) of arbitrary list %v", toString(tc.list), arbitraryBase.startPosition, toString(arbitraryBase.list)), func(t *testing.T) {
			base := New(arbitraryBase.list...)
			base.InsertAt(arbitraryBase.startPosition, tc.list...)
			got := base.String()
			assertEqual(t, got, tc.onArbitrary.onStartPosition)
		})

		t.Run(fmt.Sprintf("%v into middle position (%d) of arbitrary list %v", toString(tc.list), arbitraryBase.middlePosition, toString(arbitraryBase.list)), func(t *testing.T) {
			base := New(arbitraryBase.list...)
			base.InsertAt(arbitraryBase.middlePosition, tc.list...)
			got := base.String()
			assertEqual(t, got, tc.onArbitrary.onMiddlePosition)
		})

		t.Run(fmt.Sprintf("%v into end position (%d) of arbitrary list %v", toString(tc.list), arbitraryBase.endPosition, toString(arbitraryBase.list)), func(t *testing.T) {
			base := New(arbitraryBase.list...)
			base.InsertAt(arbitraryBase.endPosition, tc.list...)
			got := base.String()
			assertEqual(t, got, tc.onArbitrary.onEndPosition)
		})

		t.Run(fmt.Sprintf("%v into error position (%d) of arbitrary list %v", toString(tc.list), arbitraryBase.errorPosition, toString(arbitraryBase.list)), func(t *testing.T) {
			base := New(arbitraryBase.list...)
			base.InsertAt(arbitraryBase.errorPosition, tc.list...)
			got := base.String()
			assertEqual(t, got, tc.onArbitrary.onErrorPosition)
		})
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
		assertEqual(t, got, want)
	})

	type result struct {
		value int
		ok    bool
	}

	testCases := []struct {
		index int
		want  result
	}{
		{0, result{arbitrary[0], true}},
		{1, result{arbitrary[1], true}},
		{2, result{arbitrary[2], true}},
		{3, result{0, false}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("index %d on arbitrary list %v", tc.index, arbitrary), func(t *testing.T) {
			arbitrary := New(arbitrary...)
			value, ok := arbitrary.Remove(tc.index)
			assertEqual(t, value, tc.want.value)
			assertEqual(t, ok, tc.want.ok)
		})
	}
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
		{[][]int{}, toString(emptyBase), toString(arbitraryBase)},
		{[][]int{emptyList, arbitraryList}, toString(arbitraryList), toString(append(arbitraryBase, arbitraryList...))},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v on empty list %v", tc.input, emptyBase), func(t *testing.T) {
			list := New(emptyBase...)
			input := helpers.Apply(tc.input, func(l []int) *List[int] { return New(l...) })
			list.Concat(input...)
			assertEqual(t, toString(list), tc.onEmpty)
		})

		t.Run(fmt.Sprintf("%v on arbitrary list %v", tc.input, arbitraryBase), func(t *testing.T) {
			list := New(arbitraryBase...)
			input := helpers.Apply(tc.input, func(l []int) *List[int] { return New(l...) })
			list.Concat(input...)
			assertEqual(t, toString(list), tc.onArbitrary)
		})
	}
}

/* -------------------------------------------------------------------------- */

func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func toString[T any](values T) string {
	return fmt.Sprintf("%v", values)
}
