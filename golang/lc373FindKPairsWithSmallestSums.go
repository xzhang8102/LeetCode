package golang

import "container/heap"

/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n, m, ans := len(nums1), len(nums2), [][]int{}
	flag := n > m
	// make len(nums1) <= len(nums2)
	if flag {
		n, m, nums1, nums2 = m, n, nums2, nums1
	}
	if n > k {
		n = k
	}
	pq := make(lc373Heap, n)
	for i := 0; i < n; i++ {
		pq[i] = []int{nums1[i] + nums2[0], i, 0}
	}
	heap.Init(&pq)
	for pq.Len() > 0 && len(ans) < k {
		poll := heap.Pop(&pq).([]int)
		a, b := poll[1], poll[2]
		if flag {
			ans = append(ans, []int{nums2[b], nums1[a]})
		} else {
			ans = append(ans, []int{nums1[a], nums2[b]})
		}
		if b < m-1 {
			heap.Push(&pq, []int{nums1[a] + nums2[b+1], a, b + 1})
		}
	}
	return ans
}

// 最小堆模板
type lc373Heap [][]int

func (h lc373Heap) Len() int            { return len(h) }
func (h lc373Heap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h lc373Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *lc373Heap) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *lc373Heap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// @lc code=end
