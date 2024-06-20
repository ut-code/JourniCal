package helper

import (
	"sync"
)

// Thread-safe map
type Map[K comparable, V any] struct {
	mutex        sync.Mutex
	internal_map map[K]V
}

func (m *Map[K, V]) Get(k K) (v V, ok bool) {
	m.mutex.Lock()
	v, ok = m.internal_map[k]
	m.mutex.Unlock()
	return
}

func (m *Map[K, V]) Set(k K, v V) {
	m.mutex.Lock()
	m.internal_map[k] = v
	m.mutex.Unlock()
}

func (m *Map[K, V]) UnSet(k K) {
	m.mutex.Lock()
	delete(m.internal_map, k)
	m.mutex.Unlock()
}

func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{
		mutex:        sync.Mutex{},
		internal_map: make(map[K]V),
	}
}
