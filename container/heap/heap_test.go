package heap

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func checkHeap(h Interface) bool {
	for i := 0; i < h.Len(); i++ {
		left := 2*i + 1
		if left < h.Len() && h.Less(left, i) {
			return false
		}
		right := left + 1
		if right < h.Len() && h.Less(right, i) {
			return false
		}
	}
	return true
}

func TestIntHeap(t *testing.T) {
	assert := require.New(t)
	heap := NewIntHeap()
	Init(heap)
	ok := checkHeap(heap)
	assert.True(ok)

	const round = 1024 * 1024
	for i := 0; i < round; i++ {
		op := rand.Intn(3)
		switch op {
		case 0:
			heap.Push(rand.Int())
		case 1:
			if heap.Len() > 0 {
				heap.Pop()
			}
		case 2:
			if heap.Len() > 0 {
				j := rand.Intn(heap.Len())
				_, ok := heap.Remove(j)
				assert.True(ok)
			}
		}
		checkHeap(heap)
	}
}
