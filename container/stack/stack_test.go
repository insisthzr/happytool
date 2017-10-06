package stack

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type StackOp int

const (
	push StackOp = iota
	pop
)

const (
	stackOpStart = 0
	stackOpEnd   = 1
)

type StackOpRand struct {
	r *rand.Rand
}

func (r *StackOpRand) Rand() StackOp {
	len := stackOpEnd - stackOpStart + 1
	n := r.r.Intn(len)
	return StackOp(n - stackOpStart)
}

func NewStackOpRand(seed int64) *StackOpRand {
	return &StackOpRand{
		r: rand.New(rand.NewSource(seed)),
	}
}

func TestStack(t *testing.T) {
	const numRound = 1024 * 1024
	assert := require.New(t)
	now := time.Now().Unix()
	t.Log("seed", now)
	opRand := NewStackOpRand(now)
	nRand := rand.New(rand.NewSource(now))
	trustStack := []int{}
	stack := NewStack()

	for i := 0; i < numRound; i++ {
		op := opRand.Rand()
		switch op {
		case push:
			elem := nRand.Int()
			stack.Push(elem)
			trustStack = append(trustStack, elem)
		case pop:
			elem1, ok1 := stack.Pop()
			var elem2 interface{}
			ok2 := false
			length := len(trustStack)
			if length > 0 {
				ok2 = true
				elem2 = trustStack[length-1]
				trustStack = trustStack[:length-1]
			}
			assert.Equal(ok2, ok1)
			assert.Equal(elem2, elem1)
		default:
			panic("unreachable")
		}
	}

	stack = NewStack()
	objs := []interface{}{1, "2", 3.0}
	stack.Push(objs[0])
	stack.Push(objs[1])
	stack.Push(objs[2])
	i := 0
	stack.For(func(val interface{}) bool {
		assert.Equal(objs[i], val)
		i++
		return true
	})
}

func TestStringStack(t *testing.T) {
	assert := require.New(t)
	stack := NewStringStack()

	stack.Push("1")
	str, ok := stack.Pop()
	assert.True(ok)
	assert.Equal("1", str)
	str, ok = stack.Pop()
	assert.False(ok)
	val, ok := stack.Top()
	assert.False(ok)
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	stack.Push("4")
	val, ok = stack.Top()
	assert.True(ok)
	assert.Equal("4", val)
	l := stack.Len()
	assert.Equal(4, l)
	var i int64 = 1
	stack.For(func(val string) bool {
		assert.Equal(strconv.FormatInt(i, 10), val)
		i++
		return true
	})

	stack = NewStringStackWithConfig(Config{
		Capacity: 1024,
	})
	assert.Equal(1024, cap(stack.stack.data))
}
