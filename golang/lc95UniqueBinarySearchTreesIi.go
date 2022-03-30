package golang

/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	var backtrack func(start, end int) []*TreeNode
	backtrack = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		trees := []*TreeNode{}
		for i := start; i <= end; i++ {
			left := backtrack(start, i-1)
			right := backtrack(i+1, end)
			for _, leafL := range left {
				for _, leafR := range right {
					root := &TreeNode{i, leafL, leafR}
					trees = append(trees, root)
				}
			}
		}
		return trees
	}
	return backtrack(1, n)
}

// @lc code=end
