package golang

/*
 * @lc app=leetcode.cn id=801 lang=golang
 *
 * [801] Minimum Swaps To Make Sequences Increasing
 */

// @lc code=start
func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	// dp[i][0] how many swaps so far if we do not swap at index i
	// dp[i][1] how many swaps so far if we swap at index i
	dp := make([][2]int, n)
	dp[0][1] = 1
	for i := 1; i < n; i++ {
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
				// does not matter if swapped
				dp[i][0] = lc801Min(dp[i-1][0], dp[i-1][1])
				dp[i][1] = dp[i][0] + 1
			} else {
				dp[i][0] = dp[i-1][0]
				dp[i][1] = dp[i-1][1] + 1
			}
		} else if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			dp[i][0] = dp[i-1][1]
			dp[i][1] = dp[i-1][0] + 1
		}
	}
	return lc801Min(dp[n-1][0], dp[n-1][1])
}

func lc801Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
