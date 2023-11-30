package hashmap

import "fmt"

const hashMap = "HashMap"

type Map[K comparable, V any] struct {
	m map[K]V
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{m: make(map[K]V)}
}

func (m *Map[K, V]) Name() string {
	return hashMap
}

func (m *Map[K, V]) Size() int {
	return len(m.m)
}

func (m *Map[K, V]) Empty() bool {
	return len(m.m) == 0
}

func (m *Map[K, V]) String() string {
	result := ""
	for k, v := range m.m {
		result += fmt.Sprintf("%v: %v\n", k, v)
	}
	return result
}

func (m *Map[K, V]) Reset() {
	m.m = make(map[K]V)
}

func (m *Map[K, V]) Put(key K, value V) {
	m.m[key] = value
}

func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	value, ok = m.m[key]
	return
}

func (m *Map[K, V]) Remove(key K) (ok bool) {
	if _, ok = m.m[key]; ok {
		delete(m.m, key)
		return
	}
	return
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, m.Size())
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, m.Size())
	for _, v := range m.m {
		values = append(values, v)
	}
	return values
}
