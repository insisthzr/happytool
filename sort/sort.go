package sort

type Interface interface {
	Len() int
	Less(i int, j int) bool
	Swap(i int, j int)
}

func Sort(data Interface) {
	quickSort(data, 0, data.Len()-1)
}

func quickSort(data Interface, low int, high int) {
	if low >= high {
		return
	}
	mid := partition(data, low, high)
	quickSort(data, low, mid-1)
	quickSort(data, mid+1, high)
}

func partition(data Interface, low int, high int) int {
	pivot := high
	i := low - 1
	j := low
	for ; j < high; j++ {
		if data.Less(j, pivot) {
			i++
			data.Swap(j, i)
		}
	}
	if data.Less(high, i+1) {
		data.Swap(high, i+1)
	}
	return i + 1
}

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p IntSlice) Sort() { Sort(p) }

func Ints(a []int) {
	Sort(IntSlice(a))
}
