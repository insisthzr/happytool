package table

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntTable(t *testing.T) {
	assert := require.New(t)
	table := NewIntTable()

	ok := table.Set(1, 1)
	assert.True(ok)
	val, ok := table.Get(1)
	assert.True(ok)
	assert.Equal(1, val)
	val, ok = table.Get(2)
	assert.False(ok)
	assert.Nil(val)

	ok = table.Set(1, 2)
	assert.False(ok)
	val, ok = table.Get(1)
	assert.True(ok)
	assert.Equal(2, val)

	ok = table.Set(17, 3)
	assert.True(ok)
	val, ok = table.Get(1)
	assert.True(ok)
	assert.Equal(2, val)
	val, ok = table.Get(17)
	assert.True(ok)
	assert.Equal(3, val)

	ok = table.Delete(1)
	assert.True(ok)
	_, ok = table.Get(1)
	assert.False(ok)
	ok = table.Delete(4)
	assert.False(ok)
	ok = table.Set(4, 4)
	assert.True(ok)
	ok = table.Delete(4)
	assert.True(ok)
	_, ok = table.Get(4)
	assert.False(ok)
}

func TestStringTable(t *testing.T) {
	assert := require.New(t)
	table := NewStringTable()

	ok := table.Set("hello", "world")
	assert.True(ok)
	val, ok := table.Get("hello")
	assert.True(ok)
	assert.Equal("world", val)

	ok = table.Set("hello1", "world")
	assert.True(ok)
	val, ok = table.Get("hello1")
	assert.True(ok)
	assert.Equal("world", val)

	ok = table.Set("hello2", "world")
	assert.True(ok)
	val, ok = table.Get("hello2")
	assert.True(ok)
	assert.Equal("world", val)

	ok = table.Set("hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽", "world")
	assert.True(ok)
	val, ok = table.Get("hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽hellohellohellohellohellohello哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽哈喽")
	assert.True(ok)
	assert.Equal("world", val)

	ok = table.Delete("hello")
	assert.True(ok)
	_, ok = table.Get("hello")
	assert.False(ok)
	ok = table.Delete("hello")
	assert.False(ok)
}
