package main

import (
	"github.com/mhrdini/godsa/problems/bbgci/1_twopointers/inwardtraversal"
)

func main() {
	arr := []int{-5, -2, 3, 4, 6}
	target := 7
	inwardtraversal.PairSumSorted(arr, target)
}
