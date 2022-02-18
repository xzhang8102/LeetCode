package golang

/*
 * @lc app=leetcode.cn id=1791 lang=golang
 *
 * [1791] Find Center of Star Graph
 */

// @lc code=start
func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else {
		return edges[0][1]
	}
}

// @lc code=end
