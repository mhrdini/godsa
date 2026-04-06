package main

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
	"github.com/mhrdini/godsa/problems/bbgci/3_linkedlists/restructuring"
)

func main() {
	l := singlylinkedlist.New(1, 2, 4, 7, 3)
	restructuring.RemoveKthLastNode(l, 2)
	fmt.Println(l)
}
