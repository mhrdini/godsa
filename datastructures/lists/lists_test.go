package lists

import (
	"fmt"
	"testing"

	"github.com/mhrdini/godsa/datastructures/lists/arraylist"
	"github.com/mhrdini/godsa/datastructures/lists/doublylinkedlist"
	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
	"github.com/mhrdini/godsa/datastructures/utils"
	"github.com/mhrdini/godsa/helpers"
)

func lists[T any](base ...T) []List[T] {
	return []List[T]{
		arraylist.New(base...),
		singlylinkedlist.New(base...),
		doublylinkedlist.New(base...),
	}
}

func TestReset(t *testing.T) {
	arbitrary := []int{1, 2, 3}
	for _, list := range lists(arbitrary...) {
		t.Run(list.Name(), func(t *testing.T) {
			list.Reset()
			helpers.AssertEqual(t, list.Empty(), true)
		})
	}
}

func TestSort(t *testing.T) {

	orderedConstraintIntList := []int{90, 3, 29}
	orderedConstraintStringList := []string{"z", "aba", "xy"}

	for _, list := range lists(orderedConstraintIntList...) {
		t.Run(fmt.Sprintf("using OrderedComparator for int %v", list.Name()), func(t *testing.T) {
			list.Sort(utils.OrderedComparator[int])
			got := list.String()
			want := helpers.ToString([]int{3, 29, 90})
			helpers.AssertEqual(t, got, want)
		})
	}

	for _, list := range lists(orderedConstraintStringList...) {
		t.Run(fmt.Sprintf("using OrderedComparator for string %v", list.Name()), func(t *testing.T) {
			list.Sort(utils.OrderedComparator[string])
			got := list.String()
			want := helpers.ToString([]string{"aba", "xy", "z"})
			helpers.AssertEqual(t, got, want)
		})
	}

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

	movies := []movie{movie1, movie2, movie3}

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
		for _, list := range lists(movies...) {
			t.Run(fmt.Sprintf("%v on %v", tc.name, list.Name()), func(t *testing.T) {
				list.Sort(tc.comp)
				got := list.String()
				helpers.AssertEqual(t, got, tc.want)
			})
		}
	}
}

func TestAdd(t *testing.T) {

	empty := []int{}
	arbitrary := []int{1, 2, 3}

	testCases := []struct {
		base  []int
		input []int
		want  string
	}{
		{empty, empty, helpers.ToString(empty)},
		{empty, arbitrary, helpers.ToString(arbitrary)},
		{arbitrary, empty, helpers.ToString(arbitrary)},
		{arbitrary, arbitrary, helpers.ToString(append(arbitrary, arbitrary...))},
	}

	checkAddResult := func(t testing.TB, list List[int], input []int, want string) {
		t.Helper()
		list.Add(input...)
		got := list.String()
		helpers.AssertEqual(t, got, want)
	}

	for _, tc := range testCases {
		for _, list := range lists(tc.base...) {
			t.Run(fmt.Sprintf("to %v %v <- %v", list.Name(), tc.base, tc.input), func(t *testing.T) {
				checkAddResult(t, list, tc.input, tc.want)
			})
		}
	}
}

func TestInsertAt(t *testing.T) {

	inserted := []int{0, 0, 0}

	checkInserted := func(t testing.TB, list List[int], index int, want bool) {
		got := list.InsertAt(index, inserted...)
		helpers.AssertEqual(t, got, want)
	}

	checkInsertResult := func(t testing.TB, list List[int], index int, want string) {
		list.InsertAt(index, inserted...)
		got := list.String()
		helpers.AssertEqual(t, got, want)
	}

	emptyTestCases := []struct {
		index int
		want  bool
	}{
		{0, true},
		{1, false},
	}

	for _, tc := range emptyTestCases {
		for _, list := range lists[int]() {
			t.Run(fmt.Sprintf("to empty %v at index %d", list.Name(), tc.index), func(t *testing.T) {
				checkInserted(t, list, tc.index, tc.want)
			})
		}
	}

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
		for _, list := range lists(arbitrary...) {
			t.Run(fmt.Sprintf("to %v %v <- %v at index %d", list.Name(), arbitrary, inserted, tc.index), func(t *testing.T) {
				checkInsertResult(t, list, tc.index, tc.want)
			})
		}
	}
}

func TestRemove(t *testing.T) {

	empty := []int{}
	arbitrary := []int{1, 2, 3}

	for _, list := range lists(empty...) {
		t.Run(fmt.Sprintf("any index on empty %v", list.Name()), func(t *testing.T) {
			_, ok := list.Remove(0)
			got := ok
			want := false
			helpers.AssertEqual(t, got, want)
		})
	}

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
		for _, list := range lists(arbitrary...) {
			t.Run(fmt.Sprintf("index %d on arbitrary %v %v", tc.index, list.Name(), arbitrary), func(t *testing.T) {
				value, ok := list.Remove(tc.index)
				helpers.AssertEqual(t, value, tc.want.value)
				helpers.AssertEqual(t, ok, tc.want.ok)
				helpers.AssertEqual(t, list.String(), tc.want.list)
			})
		}
	}
}

func TestGet(t *testing.T) {
	empty := []int{}

	for _, list := range lists(empty...) {
		t.Run(fmt.Sprintf("any input on empty %v", list.Name()), func(t *testing.T) {
			want := false
			_, ok := list.Get(0)
			helpers.AssertEqual(t, ok, want)
			_, ok = list.Get(1)
			helpers.AssertEqual(t, ok, want)
		})
	}

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
		for _, list := range lists(arbitrary...) {
			t.Run(fmt.Sprintf("get index %d on arbitrary %v %v", tc.index, list.Name(), arbitrary), func(t *testing.T) {
				value, ok := list.Get(tc.index)
				helpers.AssertEqual(t, value, tc.want.value)
				helpers.AssertEqual(t, ok, tc.want.ok)
			})
		}
	}
}

func TestSet(t *testing.T) {

	empty := []int{}

	for _, list := range lists(empty...) {
		t.Run(fmt.Sprintf("any input on empty %v", list.Name()), func(t *testing.T) {
			want := helpers.ToString(empty)

			list.Set(0, 0)
			got := list.String()
			helpers.AssertEqual(t, got, want)

			list.Set(1, 0)
			got = list.String()
			helpers.AssertEqual(t, got, want)
		})
	}

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
		for _, list := range lists(arbitrary...) {
			t.Run(fmt.Sprintf("%d into index %d in arbitrary %v %v", tc.value, tc.index, list.Name(), arbitrary), func(t *testing.T) {
				list.Set(tc.index, tc.value)
				got := list.String()
				want := tc.want
				helpers.AssertEqual(t, got, want)
			})
		}
	}
}
