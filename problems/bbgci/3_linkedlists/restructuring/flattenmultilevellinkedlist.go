package restructuring

import "fmt"

/**
 * Problem:
 * In a multi-level linked list, each node has a next pointer
 * and child pointer. The next pointer connects to the subsequent node
 * in the same linked list, while the child pointer points to the head
 * of a new linked list under it. This creates multiple levels of
 * linked lists. If a node does not have a child list, its child attribute
 * is set to null.
 *
 * Flatten the multi-level linked list into a single-level linked list
 * by linking the end of each level to the start of the next one.
 */

type List[T comparable] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

type Node[T comparable] struct {
	value T
	child *Node[T]
	next  *Node[T]
}

func (n *Node[T]) String() string {
	if n == nil {
		return "nil"
	}

	if n.child != nil {
		return fmt.Sprintf("(%v, %v)", n.value, n.child.value)
	}

	return fmt.Sprintf("(%v)", n.value)
}

func (l *List[T]) String() string {
	vs := []*Node[T]{}
	var childNodes []*Node[T]
	visited := make(map[*Node[T]]struct{})

	for curr := l.head; curr != nil; {
		vs = append(vs, curr)
		visited[curr] = struct{}{}
		if curr.child != nil {
			childNodes = append(childNodes, curr.child)
		}
		curr = curr.next
		for curr == nil && len(childNodes) > 0 {
			curr = childNodes[0]
			childNodes = childNodes[1:]
			if _, ok := visited[curr]; ok {
				curr = nil
			}
		}
	}
	return fmt.Sprintf("%v", vs)
}

func FlattenMultiLevelLinkedList[T comparable](l *List[T]) {
	if l == nil || l.head == nil {
		return
	}

	var childNodes []*Node[T]
	visited := make(map[*Node[T]]struct{})

	curr := l.head
	var prev *Node[T]

	for curr != nil {
		visited[curr] = struct{}{}

		// queue child
		if curr.child != nil {
			childNodes = append(childNodes, curr.child)
			curr.child = nil
		}

		prev = curr

		// move forward if possible
		if curr.next != nil {
			curr = curr.next
			continue
		}

		// otherwise pull from queue
		curr = nil
		for len(childNodes) > 0 {
			next := childNodes[0]
			childNodes = childNodes[1:]

			if _, seen := visited[next]; !seen {
				curr = next
				break
			}
		}

		if curr != nil {
			prev.next = curr
		}
	}

	l.tail = prev
}

func NewMultiLevelLinkedList[T comparable](
	nextChains [][]T,
	parentToChild map[T]T,
) *List[T] {

	l := &List[T]{}
	nodes := make(map[T]*Node[T])

	// Step 1: create nodes from nextChains
	for _, chain := range nextChains {
		for _, v := range chain {
			if _, ok := nodes[v]; !ok {
				nodes[v] = &Node[T]{value: v}
				l.size++
			}
		}
	}

	// Step 2: create nodes from parentToChild (in case some aren't in chains)
	for p, c := range parentToChild {
		if _, ok := nodes[p]; !ok {
			nodes[p] = &Node[T]{value: p}
			l.size++
		}
		if _, ok := nodes[c]; !ok {
			nodes[c] = &Node[T]{value: c}
			l.size++
		}
	}

	// Step 3: connect next chains
	for i, chain := range nextChains {
		var prev *Node[T]

		for j, v := range chain {
			node := nodes[v]

			if i == 0 && j == 0 {
				l.head = node
			} else if prev != nil {
				prev.next = node
			}

			prev = node
		}

		if i == 0 {
			l.tail = prev
		}
	}

	// Step 4: connect parent → child
	for p, c := range parentToChild {
		nodes[p].child = nodes[c]
	}

	return l
}

func FlattenMultiLevelLinkedListExample() {
	levels := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9}}
	parentToChild := map[int]int{
		2: 6,
		4: 8,
		7: 10,
		8: 11,
	}
	l := NewMultiLevelLinkedList(levels, parentToChild)
	fmt.Println(l)
	FlattenMultiLevelLinkedList(l)
	fmt.Println(l)
}
