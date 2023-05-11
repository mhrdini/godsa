package avltree

import (
	"testing"

	"github.com/mhrdini/godsa/datastructures/utils/comparator"
	"github.com/mhrdini/godsa/helpers"
)

var compare = comparator.OrderedComparator[int]

func TestNew(t *testing.T) {
	tree := New(compare)

	got := tree.String()
	want := "[]"
	helpers.AssertEqual(t, got, want)

	tree = New(compare, 3, 1, 2, 2)
	got = tree.String()
	want = "[1 2 3]"
	helpers.AssertEqual(t, got, want)
}

func TestInsert(t *testing.T) {
	tree := New(compare)
	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(1)
	tree.Insert(3)
	got := tree.String()
	want := "[1 3 8]"
	helpers.AssertEqual(t, got, want)
}
