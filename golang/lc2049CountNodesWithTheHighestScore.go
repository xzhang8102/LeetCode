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
	childNum := make([]int, n)
	var dfs func(parent int)
	dfs = func(parent int) {
		if childNum[parent] != 0 {
			return
		}
		if childMap[parent] == nil {
			childNum[parent] = 1
		} else {
			childNum[parent] = 1
			for _, child := range childMap[parent] {
				dfs(child)
				childNum[parent] += childNum[child]
			}
		}
	}
	dfs(0)
	for i := 0; i < n; i++ {
		product := 1
		if i > 0 {
			product = n - childNum[i]
		}
		for _, child := range childMap[i] {
			product *= childNum[child]
		}
		if i == 0 || product > max {
			ans = 1
			max = product
		} else if product == max {
			ans++
		}
	}
	return ans
}

// @lc code=end
