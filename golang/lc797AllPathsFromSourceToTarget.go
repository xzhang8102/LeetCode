package golang

/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] All Paths From Source to Target
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	ans := [][]int{}
	row := len(graph)
	if row == 0 {
		return ans
	}
	var dfs func(node int)
	path := []int{}
	dfs = func(node int) {
		path = append(path, node)
		if node == row-1 {
			tmp := []int{}
			tmp = append(tmp, path...)
			ans = append(ans, tmp)
			return
		}
		for _, next := range graph[node] {
			dfs(next)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

// @lc code=end
