package golang

/*
 * @lc app=leetcode.cn id=2049 lang=golang
 *
 * [2049] Count Nodes With the Highest Score
 */

// @lc code=start
func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	childMap := map[int][]int{}
	for i := 1; i < n; i++ {
		if childMap[parents[i]] == nil {
			childMap[parents[i]] = []int{i}
		} else {
			childMap[parents[i]] = append(childMap[parents[i]], i)
		}
	}
	ans := 0
	max := -1
	var dfs func(parent int) int
	dfs = func(parent int) int {
		score, size := 1, 1
		for _, child := range childMap[parent] {
			cut := dfs(child)
			score *= cut
			size += cut
		}
		if size < n {
			score *= n - size
		}
		if score == max {
			ans++
		} else if score > max {
			ans = 1
			max = score
		}
		return size
	}
	dfs(0)
	return ans
}

// @lc code=end
