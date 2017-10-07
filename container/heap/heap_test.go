package heap

import (
	"math/rand"
	"testing"
	"time"

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
	seed := time.Now().Unix()
	t.Log("seed:", seed)
	r := rand.New(rand.NewSource(seed))
	heap := NewIntHeap()
	Init(heap)
	ok := checkHeap(heap)
	assert.True(ok)

	const round = 1024 * 1024
	for i := 0; i < round; i++ {
		op := r.Intn(3)
		switch op {
		case 0:
			Push(heap, r.Int())
		case 1:
			if heap.Len() > 0 {
				Pop(heap)
			}
		case 2:
			if heap.Len() > 0 {
				j := r.Intn(heap.Len())
				Remove(heap, j)
			}
		}
		assert.True(checkHeap(heap))
	}
}
