package golang

import "container/heap"

/*
 * @lc app=leetcode.cn id=1405 lang=golang
 *
 * [1405] Longest Happy String
 */

// @lc code=start
type charPair struct {
	char byte
	cnt  int
}

type lc1405PQ []charPair

func (pq *lc1405PQ) Len() int           { return len(*pq) }
func (pq *lc1405PQ) Less(i, j int) bool { return (*pq)[i].cnt > (*pq)[j].cnt }
func (pq *lc1405PQ) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
func (pq *lc1405PQ) Pop() interface{} {
	item := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return item
}
func (pq *lc1405PQ) Push(x interface{}) {
	*pq = append(*pq, x.(charPair))
}

func longestDiverseString(a int, b int, c int) string {
	var pq lc1405PQ
	pq = append(pq, charPair{'a', a})
	pq = append(pq, charPair{'b', b})
	pq = append(pq, charPair{'c', c})
	heap.Init(&pq)
	ans := []byte{}
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(charPair)
		n := len(ans)
		if n >= 2 && ans[n-2] == curr.char && ans[n-1] == curr.char {
			if pq.Len() == 0 {
				break
			}
			next := heap.Pop(&pq).(charPair)
			if next.cnt > 0 {
				ans = append(ans, next.char)
				next.cnt--
				heap.Push(&pq, next)
			}
			heap.Push(&pq, curr)
		} else {
			if curr.cnt > 0 {
				ans = append(ans, curr.char)
				curr.cnt--
				heap.Push(&pq, curr)
			}
		}
	}
	return string(ans)
}

// @lc code=end
