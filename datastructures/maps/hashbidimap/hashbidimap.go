package hashbidimap

import "fmt"

const hashBidiMap = "HashBidirectionalMap"

type Map[K comparable, V comparable] struct {
	original map[K]V
	inverse  map[V]K
}

func New[K comparable, V comparable]() *Map[K, V] {
	return &Map[K, V]{original: make(map[K]V), inverse: make(map[V]K)}
}

func (m *Map[K, V]) Name() string {
	return hashBidiMap
}

func (m *Map[K, V]) Size() int {
	return len(m.original)
}

func (m *Map[K, V]) Empty() bool {
	return len(m.original) == 0
}

func (m *Map[K, V]) String() string {
	result := ""
	for k, v := range m.original {
		result += fmt.Sprintf("%v: %v\n", k, v)
	}
	return result
}

func (m *Map[K, V]) Reset() {
	m.original = make(map[K]V)
	m.inverse = make(map[V]K)
}

func (m *Map[K, V]) Put(key K, value V) {
	m.original[key] = value
	m.inverse[value] = key
}

func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	value, ok = m.original[key]
	return
}

func (m *Map[K, V]) GetKey(value V) (key K, ok bool) {
	key, ok = m.inverse[value]
	return
}

func (m *Map[K, V]) Remove(key K) (ok bool) {
	var value V
	if value, ok = m.original[key]; ok {
		delete(m.original, key)
		delete(m.inverse, value)
		return
	}
	return
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, m.Size())
	for v := range m.inverse {
		values = append(values, v)
	}
	return values
}
