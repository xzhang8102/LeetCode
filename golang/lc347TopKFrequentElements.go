package golang

import "container/heap"

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
var freq map[int]int // value: frequency
type pq []int

func (p pq) Len() int           { return len(p) }
func (p pq) Less(i, j int) bool { return freq[p[i]] > freq[p[j]] }
func (p pq) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *pq) Push(x interface{}) {
	*p = append(*p, x.(int))
}
func (p *pq) Pop() interface{} {
	old := *p
	x := old[len(old)-1]
	*p = old[:len(old)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	ans := make([]int, 0)
	stats := make(map[int]int)
	for _, v := range nums {
		stats[v]++
	}
	freq = stats
	p := make(pq, 0)
	heap.Init(&p)
	for key := range freq {
		heap.Push(&p, key)
	}
	for i := 0; i < k; i++ {
		tmp := heap.Pop(&p)
		ans = append(ans, tmp.(int))
	}
	return ans
}

// @lc code=end
