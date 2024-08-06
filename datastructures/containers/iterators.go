package containers

type Iterator[T any] interface {
	Next() (T, bool)
	Value() (T, bool)
	Reset()
	// HasNext() bool
	// Prev() T
}

type IteratorWithIndex[T any] interface {
	Iterator[T]
	Index() int
}

type IteratorWithKey[K comparable, V any] interface {
	Iterator[V]
	Key() K
}
