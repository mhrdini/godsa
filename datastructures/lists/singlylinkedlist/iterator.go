package singlylinkedlist

import (
	"github.com/mhrdini/godsa/datastructures/containers"
)

const (
	NilIndex  = -1
	ZeroIndex = 0
)

type Iterator[T any] struct {
	list    *List[T]
	index   int
	current *node[T]
}

func (l *List[T]) NewIterator() containers.Iterator[T] {
	return &Iterator[T]{
		list:    l,
		index:   NilIndex,
		current: nil,
	}
}

func (i *Iterator[T]) Next() (result T, ok bool) {
	if i.current == nil && i.index == NilIndex && i.list.size > 0 {
		i.current = i.list.head
		i.index = ZeroIndex
		result = i.current.value
		ok = true
	} else if i.current.next != nil {
		i.current = i.current.next
		i.index++
		result = i.current.value
		ok = true
	}
	return
}

func (i *Iterator[T]) Index() int {
	return i.index
}

func (i *Iterator[T]) Value() (result T, ok bool) {
	if i.current != nil {
		result, ok = i.current.value, true
	}
	return
}
