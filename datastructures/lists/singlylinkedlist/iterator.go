package singlylinkedlist

import (
	"github.com/mhrdini/godsa/datastructures/containers"
)

const (
	NilIndex  = -1
	ZeroIndex = 0
)

type Iterator[T any] struct {
	originalList *List[T]
	list         *List[T]
	index        int
	current      *Node[T]
}

func (l *List[T]) NewIterator() containers.IteratorWithIndex[*Node[T]] {
	return &Iterator[T]{
		originalList: l,
		list:         l,
		index:        NilIndex,
		current:      nil,
	}
}

func (i *Iterator[T]) Next() (result *Node[T], ok bool) {
	if i.current == nil && i.index == NilIndex && i.list.size > 0 {
		i.current = i.list.head
		i.index = ZeroIndex
		result = i.current
		ok = true
	} else if i.current != nil && i.current.next != nil {
		i.current = i.current.next
		i.index++
		result = i.current
		ok = true
	}
	return
}

func (i *Iterator[T]) Index() int {
	return i.index
}

func (i *Iterator[T]) Value() (result *Node[T], ok bool) {
	if i.current != nil {
		result, ok = i.current, true
	}
	return
}

func (i *Iterator[T]) Reset() {
	i.list = i.originalList
	i.index = NilIndex
	i.current = nil
}
