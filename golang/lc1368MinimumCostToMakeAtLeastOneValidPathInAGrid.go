package golang

import "container/list"

/*
 * @lc app=leetcode.cn id=1368 lang=golang
 *
 * [1368] Minimum Cost to Make at Least One Valid Path in a Grid
 */

// @lc code=start
func minCost(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	mapping := map[[2]int]int{{0, 1}: 1, {0, -1}: 2, {1, 0}: 3, {-1, 0}: 4}
	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	deque := list.New()
	deque.PushBack([2]int{0, 0})
	record := make([]int, n*m)
	for i := range record {
		record[i] = -1
	}
	record[0] = 0
	for deque.Len() > 0 {
		front := deque.Front()
		deque.Remove(front)
		coord := front.Value.([2]int)
		r, c := coord[0], coord[1]
		if r == n-1 && c == m-1 {
			break
		}
		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]
			if newR >= 0 && newR < n && newC >= 0 && newC < m {
				cost := record[r*m+c]
				if grid[r][c] != mapping[dir] {
					cost += 1
				}
				if record[newR*m+newC] == -1 || record[newR*m+newC] > cost {
					record[newR*m+newC] = cost
					if cost == record[r*m+c] {
						deque.PushFront([2]int{newR, newC})
					} else {
						deque.PushBack([2]int{newR, newC})
					}
				}
			}
		}
	}
	return record[n*m-1]
}

// @lc code=end
