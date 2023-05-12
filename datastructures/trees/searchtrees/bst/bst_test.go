package bst

import (
	"testing"

	"github.com/mhrdini/godsa/datastructures/utils/comparator"
	"github.com/mhrdini/godsa/helpers"
)

var compare = comparator.OrderedComparator[int]

func TestNew(t *testing.T) {
	t.Run("empty initial", func(t *testing.T) {
		treeCase := New(compare)
		got := treeCase.String()
		want := "[]"
		helpers.AssertEqual(t, got, want)
	})

	t.Run("duplicate initial", func(t *testing.T) {
		treeCase := New(compare, 3, 1, 2, 2)
		got := treeCase.String()
		want := "[1 2 3]"
		helpers.AssertEqual(t, got, want)
	})
}

func TestInsert(t *testing.T) {
	tree := New(compare)
	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(1)
	tree.Insert(3)
	got := tree.String()
	want := helpers.ToString([]int{1, 3, 8})
	helpers.AssertEqual(t, got, want)
}

func TestRemove(t *testing.T) {
	tree := New(compare, 15, 6, 23, 4, 5, 11, 9, 10, 12, 13)
	tree.Remove(6)
	got := tree.String()
	want := helpers.ToString([]int{4, 5, 9, 10, 11, 12, 13, 15, 23})
	helpers.AssertEqual(t, got, want)
}
