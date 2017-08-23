package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQuickSort(t *testing.T) {
	assert := require.New(t)
	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	t.Logf("seed: %d", r)

	const rounds = 1024
	const maxLen = 1024
	for i := 0; i < rounds; i++ {
		ints1 := genRandomInts(r, maxLen)
		ints2 := copyInts(ints1)
		sort.Ints(ints1)
		Ints(ints2)
		assert.Equal(ints1, ints2)
	}
}

func genRandomInts(r *rand.Rand, maxLen int) []int {
	length := r.Intn(maxLen + 1)
	ints := make([]int, 0, length)
	for i := 0; i < length; i++ {
		ints = append(ints, r.Int())
	}
	return ints
}

func copyInts(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}
