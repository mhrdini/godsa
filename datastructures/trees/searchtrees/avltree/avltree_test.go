package avltree

import (
	"testing"

	"fmt"

	"github.com/mhrdini/godsa/datastructures/trees"
	"github.com/mhrdini/godsa/datastructures/utils/comparator"
	"github.com/mhrdini/godsa/helpers"
)

// Using examples from https://stackoverflow.com/q/3955680
var (
	simpleInsertBase  = []int{20, 4}
	normalInsertBase  = []int{20, 4, 26, 3, 9}
	complexInsertBase = []int{20, 4, 26, 3, 9, 2, 7, 11, 21, 30}
	simpleRemoveBase  = []int{2, 1, 4, 3, 5}
	normalRemoveBase  = []int{6, 2, 9, 1, 4, 8, 11, 3, 5, 7, 10, 12, 13}
	complexRemoveBase = []int{5, 2, 8, 1, 3, 7, 10, 4, 6, 9, 11, 12}
)

func basesForInsertion() [][]int {
	return [][]int{
		simpleInsertBase,
		normalInsertBase,
		complexInsertBase,
	}
}

func basesForRemoval() [][]int {
	return [][]int{
		simpleRemoveBase,
		normalRemoveBase,
		complexRemoveBase,
	}
}

func baseName(index int) string {
	switch index {
	case 0:
		return "simple case"
	case 1:
		return "normal case"
	case 2:
		return "complex case"
	default:
		return "unknown case"
	}
}

type result struct {
	traverser trees.Traverser[int]
	simple    string
	normal    string
	complex   string
}

func getResult(r result, index int) string {
	switch index {
	case 0:
		return r.simple
	case 1:
		return r.normal
	case 2:
		return r.complex
	default:
		return ""
	}
}

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
	testCases := []struct {
		input      int
		inOrder    result
		preOrder   result
		postOrder  result
		levelOrder result
	}{
		{
			input: 15,
			inOrder: result{
				traverser: trees.InOrder[int],
				simple:    helpers.ToString([]int{4, 15, 20}),
				normal:    helpers.ToString([]int{3, 4, 9, 15, 20, 26}),
				complex:   helpers.ToString([]int{2, 3, 4, 7, 9, 11, 15, 20, 21, 26, 30}),
			},
			preOrder: result{
				traverser: trees.PreOrder[int],
				simple:    helpers.ToString([]int{15, 4, 20}),
				normal:    helpers.ToString([]int{9, 4, 3, 20, 15, 26}),
				complex:   helpers.ToString([]int{9, 4, 3, 2, 7, 20, 11, 15, 26, 21, 30}),
			},
			postOrder: result{
				traverser: trees.PostOrder[int],
				simple:    helpers.ToString([]int{4, 20, 15}),
				normal:    helpers.ToString([]int{3, 4, 15, 26, 20, 9}),
				complex:   helpers.ToString([]int{2, 3, 7, 4, 15, 11, 21, 30, 26, 20, 9}),
			},
			levelOrder: result{
				traverser: trees.LevelOrder[int],
				simple:    helpers.ToString([]int{15, 4, 20}),
				normal:    helpers.ToString([]int{9, 4, 20, 3, 15, 26}),
				complex:   helpers.ToString([]int{9, 4, 20, 3, 7, 11, 26, 2, 15, 21, 30}),
			},
		},
		{
			input: 8,
			inOrder: result{
				traverser: trees.InOrder[int],
				simple:    helpers.ToString([]int{4, 8, 20}),
				normal:    helpers.ToString([]int{3, 4, 8, 9, 20, 26}),
				complex:   helpers.ToString([]int{2, 3, 4, 7, 8, 9, 11, 20, 21, 26, 30}),
			},
			preOrder: result{
				traverser: trees.PreOrder[int],
				simple:    helpers.ToString([]int{8, 4, 20}),
				normal:    helpers.ToString([]int{9, 4, 3, 8, 20, 26}),
				complex:   helpers.ToString([]int{9, 4, 3, 2, 7, 8, 20, 11, 26, 21, 30}),
			},
			postOrder: result{
				traverser: trees.PostOrder[int],
				simple:    helpers.ToString([]int{4, 20, 8}),
				normal:    helpers.ToString([]int{3, 8, 4, 26, 20, 9}),
				complex:   helpers.ToString([]int{2, 3, 8, 7, 4, 11, 21, 30, 26, 20, 9}),
			},
			levelOrder: result{
				traverser: trees.LevelOrder[int],
				simple:    helpers.ToString([]int{8, 4, 20}),
				normal:    helpers.ToString([]int{9, 4, 20, 3, 8, 26}),
				complex:   helpers.ToString([]int{9, 4, 20, 3, 7, 11, 26, 2, 8, 21, 30}),
			},
		},
	}

	for _, tc := range testCases {
		for i, base := range basesForInsertion() {
			baseName := baseName(i)
			treeCase := New(compare, base...)
			treeCase.Insert(tc.input)

			var got, want string

			t.Run(fmt.Sprintf("%d into %v %v inorder", tc.input, treeCase.Name(), baseName), func(t *testing.T) {
				got = helpers.ToString(trees.Traverse(treeCase, tc.inOrder.traverser))
				want = getResult(tc.inOrder, i)
				helpers.AssertEqual(t, got, want)
			})
			t.Run(fmt.Sprintf("%d into %v %v preorder", tc.input, treeCase.Name(), baseName), func(t *testing.T) {
				got = helpers.ToString(trees.Traverse(treeCase, tc.preOrder.traverser))
				want = getResult(tc.preOrder, i)
				helpers.AssertEqual(t, got, want)
			})
			t.Run(fmt.Sprintf("%d into %v %v postorder", tc.input, treeCase.Name(), baseName), func(t *testing.T) {
				got = helpers.ToString(trees.Traverse(treeCase, tc.postOrder.traverser))
				want = getResult(tc.postOrder, i)
				helpers.AssertEqual(t, got, want)
			})
			t.Run(fmt.Sprintf("%d into %v %v levelorder", tc.input, treeCase.Name(), baseName), func(t *testing.T) {
				got = helpers.ToString(trees.Traverse(treeCase, tc.levelOrder.traverser))
				want = getResult(tc.levelOrder, i)
				helpers.AssertEqual(t, got, want)
			})
		}
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		input      int
		inOrder    result
		preOrder   result
		postOrder  result
		levelOrder result
	}{
		{
			input: 1,
			inOrder: result{
				traverser: trees.InOrder[int],
				simple:    helpers.ToString([]int{2, 3, 4, 5}),
				normal:    helpers.ToString([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}),
				complex:   helpers.ToString([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
			},
			preOrder: result{
				traverser: trees.PreOrder[int],
				simple:    helpers.ToString([]int{4, 2, 3, 5}),
				normal:    helpers.ToString([]int{6, 4, 2, 3, 5, 9, 8, 7, 11, 10, 12, 13}),
				complex:   helpers.ToString([]int{8, 5, 3, 2, 4, 7, 6, 10, 9, 11, 12}),
			},
			postOrder: result{
				traverser: trees.PostOrder[int],
				simple:    helpers.ToString([]int{3, 2, 5, 4}),
				normal:    helpers.ToString([]int{3, 2, 5, 4, 7, 8, 10, 13, 12, 11, 9, 6}),
				complex:   helpers.ToString([]int{2, 4, 3, 6, 7, 5, 9, 12, 11, 10, 8}),
			},
			levelOrder: result{
				traverser: trees.LevelOrder[int],
				simple:    helpers.ToString([]int{4, 2, 5, 3}),
				normal:    helpers.ToString([]int{6, 4, 9, 2, 5, 8, 11, 3, 7, 10, 12, 13}),
				complex:   helpers.ToString([]int{8, 5, 10, 3, 7, 9, 11, 2, 4, 6, 12}),
			},
		},
	}

	for _, tc := range testCases {
		for i, base := range basesForRemoval() {
			baseName := baseName(i)
			treeCase := New(compare, base...)
			treeCase.Remove(tc.input)

			var got, want string

			t.Run(fmt.Sprintf("%d from %v %v inorder", tc.input, treeCase.Name(), baseName), func(t *testing.T) {
				got = helpers.ToString(trees.Traverse(treeCase, tc.inOrder.traverser))
				want = getResult(tc.inOrder, i)
				helpers.AssertEqual(t, got, want)
			})
			t.Run(fmt.Sprintf("%d from %v %v preorder", tc.input, treeCase.Name(), baseName), func(t *testing.T) {
				got = helpers.ToString(trees.Traverse(treeCase, tc.preOrder.traverser))
				want = getResult(tc.preOrder, i)
				helpers.AssertEqual(t, got, want)
			})
			t.Run(fmt.Sprintf("%d from %v %v postorder", tc.input, treeCase.Name(), baseName), func(t *testing.T) {
				got = helpers.ToString(trees.Traverse(treeCase, tc.postOrder.traverser))
				want = getResult(tc.postOrder, i)
				helpers.AssertEqual(t, got, want)
			})
			t.Run(fmt.Sprintf("%d from %v %v levelorder", tc.input, treeCase.Name(), baseName), func(t *testing.T) {
				got = helpers.ToString(trees.Traverse(treeCase, tc.levelOrder.traverser))
				want = getResult(tc.levelOrder, i)
				helpers.AssertEqual(t, got, want)
			})
		}
	}
}
