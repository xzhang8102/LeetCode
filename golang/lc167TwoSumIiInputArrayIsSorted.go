package golang

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	ans := make([]int, 2)
	n := len(numbers)
	for lo, hi := 0, n-1; lo < hi; {
		if numbers[lo]+numbers[hi] == target {
			ans[0], ans[1] = lo+1, hi+1
			return ans
		} else if numbers[lo]+numbers[hi] > target {
			hi--
		} else {
			lo++
		}
	}
	return ans
}

// @lc code=end
