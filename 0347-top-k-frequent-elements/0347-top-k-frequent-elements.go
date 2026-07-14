var freq map[int]int

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool {
	return freq[h[i]] < freq[h[j]]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	freq = make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	h := &IntHeap{}
	heap.Init(h)

	for num := range freq {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([]int, 0, k)
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(int))
	}

	return ans
}