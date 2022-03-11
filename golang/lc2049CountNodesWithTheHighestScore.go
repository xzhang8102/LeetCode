package golang

/*
 * @lc app=leetcode.cn id=2049 lang=golang
 *
 * [2049] Count Nodes With the Highest Score
 */

// @lc code=start
func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	childNum := make([]int, n)
	childMap := map[int][]int{}
	for i := 0; i < n; i++ {
		childNum[i]++
		if i > 0 {
			for curr := i; parents[curr] != -1; {
				childNum[parents[curr]]++
				curr = parents[curr]
			}
			if childMap[parents[i]] == nil {
				childMap[parents[i]] = []int{i}
			} else {
				childMap[parents[i]] = append(childMap[parents[i]], i)
			}
		}
	}
	ans := 0
	max := -1
	for i := 0; i < n; i++ {
		product := 1
		if i > 0 {
			product *= childNum[0] - childNum[i]
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
