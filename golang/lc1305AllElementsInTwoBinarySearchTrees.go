package golang

/*
 * @lc app=leetcode.cn id=1305 lang=golang
 *
 * [1305] All Elements in Two Binary Search Trees
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
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	inorder := func(root *TreeNode) []int {
		ans := []int{}
		if root == nil {
			return ans
		}
		stack := []interface{}{root}
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node, ok := top.(*TreeNode); ok {
				if node.Right != nil {
					stack = append(stack, node.Right)
				}
				stack = append(stack, node.Val)
				if node.Left != nil {
					stack = append(stack, node.Left)
				}
			} else {
				ans = append(ans, top.(int))
			}
		}
		return ans
	}
	arr1, arr2 := inorder(root1), inorder(root2)
	n1, n2 := len(arr1), len(arr2)
	ans := make([]int, n1+n2)
	i, j, k := 0, 0, 0
	for j < n1 && k < n2 {
		if arr1[j] <= arr2[k] {
			ans[i] = arr1[j]
			j++
		} else {
			ans[i] = arr2[k]
			k++
		}
		i++
	}
	if j < n1 {
		copy(ans[i:], arr1[j:])
	}
	if k < n2 {
		copy(ans[i:], arr2[k:])
	}
	return ans
}

// @lc code=end
