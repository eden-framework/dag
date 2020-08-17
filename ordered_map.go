package dag

import "sync"

type OrderedMap struct {
	lock sync.RWMutex
	data map[interface{}]interface{}
	keys []interface{}
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		lock: sync.RWMutex{},
		data: make(map[interface{}]interface{}),
		keys: make([]interface{}, 0),
	}
}

func (m *OrderedMap) Put(key, val interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	_, exist := m.data[key]
	if !exist {
		m.keys = append(m.keys, key)
	}
	m.data[key] = val
}

func (m *OrderedMap) Get(key interface{}) (interface{}, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	v, exist := m.data[key]
	return v, exist
}

func (m *OrderedMap) PutOrGet(key interface{}, val interface{}) (interface{}, bool) {
	v, exist := m.Get(key)
	if exist {
		return v, true
	}

	m.Put(key, val)
	return val, false
}

func (m *OrderedMap) Walk(walker func(key, val interface{}) bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for _, key := range m.keys {
		if !walker(key, m.data[key]) {
			break
		}
	}
}

func (m *OrderedMap) Size() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return len(m.data)
}
