package golang

/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] Number of Provinces
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	ans := 0
	row := len(isConnected)
	if row == 0 {
		return ans
	}
	visited := map[int]bool{}
	var dfs func(city int)
	dfs = func(city int) {
		if !visited[city] {
			visited[city] = true
			for i := range isConnected[city] {
				if i != city && isConnected[city][i] == 1 {
					dfs(i)
				}
			}
		}
	}
	for i := range isConnected {
		if !visited[i] {
			ans++
			dfs(i)
		}
	}
	return ans
}

// @lc code=end
