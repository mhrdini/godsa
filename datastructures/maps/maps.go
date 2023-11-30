package maps

import "github.com/mhrdini/godsa/datastructures/containers"

type Map[K comparable, V any] interface {
	containers.Container[V]

	Put(key K, value V)
	Get(key K) (value V, ok bool)
	Remove(key K) (ok bool)
	Keys() []K
	Values() []V
}
