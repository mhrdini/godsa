package linkedhashmap

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

const linkedHashMap = "LinkedHashMap"

type Map[K comparable, V any] struct {
	m        map[K]V
	ordering *singlylinkedlist.List[K]
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m:        make(map[K]V),
		ordering: singlylinkedlist.New[K](),
	}
}

func (m *Map[K, V]) Name() string {
	return linkedHashMap
}

func (m *Map[K, V]) Size() int {
	return len(m.m)
}

func (m *Map[K, V]) Empty() bool {
	return len(m.m) == 0
}

func (m *Map[K, V]) String() string {
	result := ""
	for _, k := range m.ordering.Values() {
		result += fmt.Sprintf("%v: %v\n", k, m.m[k])
	}
	return result
}

func (m *Map[K, V]) Reset() {
	m.m = make(map[K]V)
}

func (m *Map[K, V]) Put(key K, value V) {
	if _, ok := m.m[key]; !ok {
		m.ordering.Add(key)
	}

}

func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	value, ok = m.m[key]
	return
}

func (m *Map[K, V]) Remove(key K) (ok bool) {
	if _, ok := m.m[key]; ok {
		delete(m.m, key)
		for i, k := range m.ordering.Values() {
			if k == key {
				m.ordering.Remove(i)
				break
			}
		}
	}
	return
}

func (m *Map[K, V]) Keys() []K {
	return m.ordering.Values()
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, m.Size())
	for _, k := range m.ordering.Values() {
		values = append(values, m.m[k])
	}
	return values
}
