package golang

/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		// 讨论1至i每个数作为根节点的情况
		// 左子树的个数与右子树的个数相乘
		// 由于BST的特性，左右子树的数值是连续的
		for j := 1; j <= i; j++ {
			dp[i] += dp[i-j] * dp[j-1]
		}
	}
	return dp[n]
}

// @lc code=end
