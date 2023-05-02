package doublylinkedlist

import (
	"fmt"
	"testing"

	"github.com/mhrdini/godsa/datastructures/utils"
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

func TestSort(t *testing.T) {

	orderedConstraintIntList := []int{90, 3, 29}
	orderedConstraintStringList := []string{"z", "aba", "xy"}

	t.Run("using OrderedComparator for int list", func(t *testing.T) {
		list := New(orderedConstraintIntList...)
		list.Sort(utils.OrderedComparator[int])
		got := list.String()
		want := helpers.ToString([]int{3, 29, 90})
		helpers.AssertEqual(t, got, want)
	})

	t.Run("using OrderedComparator for string list", func(t *testing.T) {
		list := New(orderedConstraintStringList...)
		list.Sort(utils.OrderedComparator[string])
		got := list.String()
		want := helpers.ToString([]string{"aba", "xy", "z"})
		helpers.AssertEqual(t, got, want)
	})

	type rating struct {
		criticValue   float64
		audienceValue float64
	}

	type movie struct {
		title string
		score rating
	}

	byTitle := func(a, b movie) int {
		return utils.OrderedComparator(a.title, b.title)
	}
	byCriticScore := func(a, b movie) int {
		return utils.OrderedComparator(a.score.criticValue, b.score.criticValue)
	}
	byAudienceScore := func(a, b movie) int {
		return utils.OrderedComparator(a.score.audienceValue, b.score.audienceValue)
	}

	movie1 := movie{"Aftersun", rating{0.90, 0.86}}
	movie2 := movie{"Parasite", rating{0.85, 0.90}}
	movie3 := movie{"John Wick: Chapter 4", rating{0.82, 0.95}}

	list := []movie{movie1, movie2, movie3}

	testCases := []struct {
		name string
		comp func(a, b movie) int
		want string
	}{
		{"by title", byTitle, helpers.ToString([]movie{movie1, movie3, movie2})},
		{"by critic score", byCriticScore, helpers.ToString([]movie{movie3, movie2, movie1})},
		{"by audience score", byAudienceScore, helpers.ToString([]movie{movie1, movie2, movie3})},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := New(list...)
			list.Sort(tc.comp)
			got := list.String()
			helpers.AssertEqual(t, got, tc.want)
		})
	}
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
			helpers.AssertEqual(t, list.Size(), tc.wantOnEmpty)
		})

		t.Run(fmt.Sprintf("%v onto arbitrary list %v", tc.values, arbitrary), func(t *testing.T) {
			list := New(arbitrary...)
			list.Add(tc.values...)
			helpers.AssertEqual(t, list.Size(), tc.wantOnArbitrary)
		})
	}
}

func TestInsertAt(t *testing.T) {

	inserted := []int{0, 0, 0}

	t.Run("zero index on empty list", func(t *testing.T) {
		list := New[int]()

		got := list.InsertAt(0, inserted...)
		want := true
		helpers.AssertEqual(t, got, want)
	})

	t.Run("non-zero on empty list", func(t *testing.T) {
		list := New[int]()

		got := list.InsertAt(1, inserted...)
		want := false
		helpers.AssertEqual(t, got, want)
	})

	arbitrary := []int{1, 2, 3, 4, 5}

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
		t.Run(fmt.Sprintf("%v into index %d of arbitrary list %v", inserted, tc.index, arbitrary), func(t *testing.T) {
			list := New(arbitrary...)
			list.InsertAt(tc.index, inserted...)
			got := list.String()
			want := tc.want
			helpers.AssertEqual(t, got, want)
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

func TestGet(t *testing.T) {
	empty := []int{}

	t.Run("any input on empty list", func(t *testing.T) {

		list := New(empty...)
		want := false

		_, ok := list.Get(0)
		helpers.AssertEqual(t, ok, want)

		_, ok = list.Get(1)
		helpers.AssertEqual(t, ok, want)
	})

	arbitrary := []int{1, 2, 3}

	type result struct {
		value int
		ok    bool
	}

	var zeroValue int

	testCases := []struct {
		index int
		want  result
	}{
		{0, result{arbitrary[0], true}},
		{1, result{arbitrary[1], true}},
		{2, result{arbitrary[2], true}},
		{3, result{zeroValue, false}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("index %d on arbitrary list %v", tc.index, arbitrary), func(t *testing.T) {
			list := New(arbitrary...)
			value, ok := list.Get(tc.index)
			helpers.AssertEqual(t, value, tc.want.value)
			helpers.AssertEqual(t, ok, tc.want.ok)
		})
	}
}

func TestSet(t *testing.T) {

	empty := []int{}

	t.Run("any input on empty list", func(t *testing.T) {
		list := New(empty...)
		want := helpers.ToString(empty)

		list.Set(0, 0)
		got := list.String()
		helpers.AssertEqual(t, got, want)

		list.Set(1, 0)
		got = list.String()
		helpers.AssertEqual(t, got, want)
	})

	arbitrary := []int{1, 2, 3}

	testCases := []struct {
		index int
		value int
		want  string
	}{
		{0, 0, helpers.ToString([]int{0, 2, 3})},
		{1, 0, helpers.ToString([]int{1, 0, 3})},
		{2, 0, helpers.ToString([]int{1, 2, 0})},
		{3, 0, helpers.ToString([]int{1, 2, 3})},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d into index %d in arbitrary list %v", tc.value, tc.index, arbitrary), func(t *testing.T) {
			list := New(arbitrary...)
			list.Set(tc.index, tc.value)
			got := list.String()
			want := tc.want
			helpers.AssertEqual(t, got, want)
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
