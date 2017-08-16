package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	assert := require.New(t)

	list := NewList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	assert.Equal(1, list.Front().Value)
	assert.Equal(3, list.Back().Value)

	list.PushFront(4)
	list.PushBack(5)
	assert.Equal(4, list.Front().Value)
	assert.Equal(5, list.Back().Value)

	value, ok := list.PopFront()
	assert.True(ok)
	assert.Equal(4, value)
	value, ok = list.PopBack()
	assert.True(ok)
	assert.Equal(5, value)

	wants := []int{1, 2, 3}
	i := 0
	list.For(func(value interface{}) bool {
		t.Logf("value: %v", value)
		assert.Equal(wants[i], value.(int))
		i++
		return true
	})

	i = 0
	for list.Len() > 0 {
		value, ok := list.PopFront()
		assert.True(ok)
		assert.Equal(wants[i], value)
		i++
	}

	_, ok = list.PopFront()
	assert.False(ok)
	_, ok = list.PopBack()
	assert.False(ok)
}

func TestElement(t *testing.T) {
	assert := require.New(t)

	list := NewList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	current := list.Front()
	assert.NotNil(current)
	assert.Equal(1, current.Value)
	current = current.Next()
	assert.NotNil(current)
	assert.Equal(2, current.Value)
	current = current.Next()
	assert.NotNil(current)
	assert.Equal(3, current.Value)

	current = list.Back()
	assert.NotNil(current)
	assert.Equal(3, current.Value)
	current = current.Prev()
	assert.NotNil(current)
	assert.Equal(2, current.Value)
	current = current.Prev()
	assert.NotNil(current)
	assert.Equal(1, current.Value)
}

func TestInsertRemove(t *testing.T) {
	assert := require.New(t)

	list := NewList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	current := list.Front()  //1
	current = current.Next() //2
	list.InsertAfter(4, current)
	assert.NotNil(current.Next())
	assert.Equal(4, current.Next().Value)
	assert.NotNil(current.Next().Next())
	assert.Equal(3, current.Next().Next().Value)

	list.InsertBefore(5, current)
	assert.NotNil(current.Prev())
	assert.Equal(5, current.Prev().Value)

	prev := current.Prev()
	next := current.Next()
	assert.NotNil(list.Remove(current))
	assert.Equal(4, prev.Next().Value)
	assert.Equal(5, next.Prev().Value)
}
