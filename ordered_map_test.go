package dag

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var m = NewOrderedMap()

func TestOrderedMap_Put(t *testing.T) {
	m.Put("foo", "bar")
	m.Put("foo1", "bar1")

	assert.Equal(t, m.Size(), 2)
}

func TestOrderedMap_Get(t *testing.T) {
	v, exist := m.Get("foo")
	assert.Equal(t, exist, true)
	assert.Equal(t, v, "bar")
}

func TestOrderedMap_PutOrGet(t *testing.T) {
	v, exist := m.PutOrGet("foo", "bar")
	assert.Equal(t, exist, true)
	assert.Equal(t, v, "bar")

	v1, exist1 := m.PutOrGet("foo2", "bar2")
	assert.Equal(t, exist1, false)
	assert.Equal(t, v1, "bar2")
}

func TestOrderedMap_Walk(t *testing.T) {
	m.Walk(func(key, val interface{}) bool {
		fmt.Println(key, val)
		return true
	})
}

func TestOrderedMap_Size(t *testing.T) {
	assert.Equal(t, m.Size(), 3)
}
