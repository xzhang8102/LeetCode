package golang

/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	graph := make([][]int, n)
	degree := make([]int, n)
	for _, edge := range edges {
		e1, e2 := edge[0], edge[1]
		graph[e1] = append(graph[e1], e2)
		graph[e2] = append(graph[e2], e1)
		degree[e1]++
		degree[e2]++
	}
	ans := []int{}
	for i, d := range degree {
		if d == 1 {
			ans = append(ans, i)
		}
	}
	for n > 2 {
		n -= len(ans)
		tmp := ans
		ans = nil
		for _, leaf := range tmp {
			for _, neighbour := range graph[leaf] {
				degree[neighbour]--
				if degree[neighbour] == 1 {
					ans = append(ans, neighbour)
				}
			}
		}
	}
	return ans
}

// @lc code=end
