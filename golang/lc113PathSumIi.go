package golang

/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] Path Sum II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	pick := []int{}
	ans := [][]int{}
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if sum+node.Val == targetSum {
				tmp := make([]int, len(pick))
				copy(tmp, pick)
				tmp = append(tmp, node.Val)
				ans = append(ans, tmp)
			}
			return
		}
		pick = append(pick, node.Val)
		if node.Left != nil {
			dfs(node.Left, sum+node.Val)
		}
		if node.Right != nil {
			dfs(node.Right, sum+node.Val)
		}
		pick = pick[:len(pick)-1]
	}
	dfs(root, 0)
	return ans
}

// @lc code=end
