package golang

/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	ans, count := 0, 0
	n := len(nums)
	for left, right := 0, 0; right < n; {
		if nums[right] == 0 {
			if count < k {
				count++
				ans = lc1004Max(ans, right-left+1)
				right++
			} else {
				if nums[left] == 0 {
					count--
				}
				left++
			}
		} else {
			ans = lc1004Max(ans, right-left+1)
			right++
		}
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
