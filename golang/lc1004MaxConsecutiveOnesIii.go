package golang

/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	var ans, left, lsum, rsum int
	for right, v := range nums {
		rsum += 1 - v
		for lsum < rsum-k {
			lsum += 1 - nums[left]
			left++
		}
		ans = lc1004Max(ans, right-left+1)
	}
	return ans
}

func lc1004Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
