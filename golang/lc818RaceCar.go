package golang

import "fmt"

/*
 * @lc app=leetcode.cn id=818 lang=golang
 *
 * [818] Race Car
 */

// @lc code=start
func racecar(target int) int {
	q := [][2]int{{0, 1}}
	visited := map[string]bool{"0_1": true}
	ans := 0
	for len(q) > 0 {
		tmp := q
		q = [][2]int{}
		for _, curr := range tmp {
			if curr[0] == target {
				return ans
			}
			nxt := [2]int{curr[0] + curr[1], curr[1] << 1}
			key := fmt.Sprintf("%d_%d", nxt[0], nxt[1])
			if !visited[key] && nxt[0] > 0 && nxt[0] < target<<1 {
				q = append(q, nxt)
				visited[key] = true
			}
			nxt[0] = curr[0]
			nxt[1] = -1
			if curr[1] < 0 {
				nxt[1] = 1
			}
			key = fmt.Sprintf("%d_%d", nxt[0], nxt[1])
			if !visited[key] && nxt[0] > 0 && nxt[0] < target<<1 {
				q = append(q, nxt)
				visited[key] = true
			}
		}
		ans++
	}
	return -1
}

// @lc code=end
