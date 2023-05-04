package arraylist

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/utils"
)

const arrayList = "ArrayList"

type List[T any] struct {
	size int
	list []T
}

var (
	growthFactor float64 = 2.0
)

func New[T any](v ...T) *List[T] {
	list := &List[T]{}
	if len(v) > 0 {
		list.Add(v...)
	}
	return list
}

func (l *List[T]) Name() string {
	return arrayList
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Empty() bool {
	return l.size == 0
}

func (l *List[T]) Values() []T {
	return l.list[:l.size:l.size]
}

func (l *List[T]) String() string {
	return fmt.Sprintf("%v", l.Values())
}

func (l *List[T]) Reset() {
	l.size = 0
	l.list = []T{}
}

func (l *List[T]) Sort(comp utils.Comparator[T]) {
	utils.Sort(l.list[:l.size], comp)
}

func (l *List[T]) Add(vs ...T) bool {
	n := len(vs)
	if n == 0 {
		return false
	}

	l.allocateFor(n)
	for _, v := range vs {
		l.list[l.size] = v
		l.size++
	}
	return true
}

func (l *List[T]) InsertAt(i int, vs ...T) bool {
	n := len(vs)
	if n == 0 || !l.withinRange(i) {
		return false
	}

	if i == l.size {
		return l.Add(vs...)
	}

	l.allocateFor(n)
	l.size += n
	copy(l.list[i+n:], l.list[i:l.size-n])
	copy(l.list[i:], vs)
	return true
}

func (l *List[T]) Remove(i int) (T, bool) {
	var value T

	if l.Empty() || i < 0 || i >= l.size {
		return value, false
	}

	value = l.list[i]
	copy(l.list[i:], l.list[i+1:l.size])
	l.size--
	return value, true
}

func (l *List[T]) Get(i int) (T, bool) {
	var value T

	if l.Empty() || !l.withinRange(i) {
		return value, false
	}

	switch i {
	case l.size:
		return value, false
	default:
		return l.list[i], true
	}
}

func (l *List[T]) Set(i int, v T) bool {
	if !l.withinRange(i) {
		return false
	}

	switch l.size {
	case 0, i:
		return false
	default:
		l.list[i] = v
		return true
	}
}

func (l *List[T]) Concat(ls ...*List[T]) {
	if len(ls) == 0 {
		return
	}

	for _, list := range ls {
		l.Add(list.Values()...)
	}
}

func (l *List[T]) allocateFor(n int) {
	if l.size+n > cap(l.list) {
		l.resize(int(growthFactor * float64(cap(l.list)+n)))
	}
}

func (l *List[T]) resize(capacity int) {
	resized := make([]T, capacity)
	copy(resized, l.list)
	l.list = resized
}

func (l *List[T]) withinRange(i int) bool {
	return i >= 0 && i <= l.size
}
