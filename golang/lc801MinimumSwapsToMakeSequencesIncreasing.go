package golang

/*
 * @lc app=leetcode.cn id=801 lang=golang
 *
 * [801] Minimum Swaps To Make Sequences Increasing
 */

// @lc code=start
import "math"

func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	// dp[i][0] 代表nums1[0->i], nums2[0->i]均为单调递增且index i不交换的最小swap数
	// dp[i][1] 代表nums1[0->i], nums2[0->i]均为单调递增且index i交换的最小swap数
	dp := make([][2]int, n)
	for i := range dp {
		dp[i][0] = math.MaxInt32
		dp[i][1] = math.MaxInt32
	}
	dp[0][0] = 0
	dp[0][1] = 1
	for i := 1; i < n; i++ {
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			dp[i][0] = dp[i-1][0]
			// 如果已经是单调递增，必须交换已确保还是单调递增(因为真实数组并没有进行变动)
			dp[i][1] = dp[i-1][1] + 1
		}
		// 交叉检查 对于局部情况交叉检查就相当于真实交换
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			dp[i][0] = min(dp[i][0], dp[i-1][1])
			dp[i][1] = min(dp[i][1], dp[i-1][0]+1)
		}
	}
	return min(dp[n-1][0], dp[n-1][1])
}

// @lc code=end
