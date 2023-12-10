package containers

type Iterator[T any] interface {
	Next() (T, bool)
	Value() (T, bool)
	// HasNext() bool
	// Prev() T
	// Reset()
}

type IteratorWithIndex[T any] interface {
	Iterator[T]
	Index() int
}

type IteratorWithKey[K comparable, V any] interface {
	Iterator[V]
	Key() K
}
