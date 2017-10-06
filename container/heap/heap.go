package heap

import (
	"sort"
)

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func Init(h Interface) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) { //if false, this element not going down, try up
			up(h, i)
		}
	}
	return h.Pop()
}

func down(h Interface, start int, n int) bool {
	cur := start
	for {
		left := 2*cur + 1
		if left >= n {
			break
		}
		if left < 0 {
			break //int overflow
		}
		child := left
		right := left + 1
		if right < n && h.Less(right, left) {
			child = right
		}
		if !h.Less(child, cur) {
			break
		}
		h.Swap(cur, child)
		cur = child
	}
	return cur > start //has called swap
}

func up(h Interface, cur int) {
	for {
		parent := (cur - 1) / 2
		if cur == parent {
			break
		}
		if !h.Less(cur, parent) {
			break
		}
		h.Swap(cur, parent)
		cur = parent
	}
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewIntHeap() *IntHeap {
	heap := make(IntHeap, 0)
	return &heap
}
