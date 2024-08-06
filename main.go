package main

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
	linkedlists "github.com/mhrdini/godsa/problems/ctci/2_linkedlists"
)

func main() {
	list := singlylinkedlist.New("a", "b", "c", "d", "e", "c")
	linkedlists.DeleteMiddleNode(list, "c")
	fmt.Println(list)
}
