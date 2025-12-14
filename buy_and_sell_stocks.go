package main

import "container/heap"

// type MinHeap []int
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
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

func maxProfitGreedy(stockPrices []int) int {
	options := &IntHeap{}
	heap.Init(options)
	currentProfit := 0
	for i := 0; i < len(stockPrices); i++ {
		if options.Len() > 0 && stockPrices[i] > (*options)[0] {
			smallestPrice := heap.Pop(options).(int)
			currentProfit += stockPrices[i] - smallestPrice
		}
		heap.Push(options, stockPrices[i])
	}
	return currentProfit
}
