package weekly288

import (
	"container/heap"
)

type priorityQ []int

func (pq priorityQ) Len() int           { return len(pq) }
func (pq priorityQ) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq priorityQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQ) Pop() interface{} {
	x := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return x
}
func (pq *priorityQ) Push(v interface{}) {
	*pq = append(*pq, v.(int))
}

func maximumProduct(nums []int, k int) int {
	pq := priorityQ(nums)
	heap.Init(&pq)
	for k > 0 {
		pq[0]++
		heap.Fix(&pq, 0)
		k--
	}
	ans := 1
	mod := int(1e9 + 7)
	for pq.Len() > 0 {
		ans = (ans * pq.Pop().(int)) % mod
	}
	return ans
}
