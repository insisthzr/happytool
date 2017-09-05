package table

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntTable(t *testing.T) {
	assert := require.New(t)
	table := NewIntTable()

	table.Set(1, 1)
	val, ok := table.Get(1)
	assert.True(ok)
	assert.Equal(1, val)
	val, ok = table.Get(2)
	assert.False(ok)
	assert.Equal(nil, val)

	table.Set(1, 2)
	val, ok = table.Get(1)
	assert.True(ok)
	assert.Equal(2, val)

	table.Set(17, 3)
	val, ok = table.Get(1)
	assert.True(ok)
	assert.Equal(2, val)
	val, ok = table.Get(17)
	assert.True(ok)
	assert.Equal(3, val)
}
