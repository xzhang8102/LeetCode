package golang

/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0) // store the index
	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			result[top] = i - top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}

// @lc code=end
