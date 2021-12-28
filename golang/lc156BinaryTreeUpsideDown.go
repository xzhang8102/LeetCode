package golang

/*
 * @lc app=leetcode.cn id=156 lang=golang
 *
 * [156] Binary Tree Upside Down
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
/*
test case:
   1
  / \
 2   3
最直观的想法就是按照题目要求把各个节点重新连接起来
newRoot = root.Left
newRoot.Left = root.Right
newRoot.Right = root
此时root应该是叶子结点，应该清理干净
root.Left = null
root.Right = null
此时root的左节点恰好是新的root，所以
return newRoot
至此这道题的递归框架基本上就出来了
*/
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	// 递归终止条件
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	newRoot := upsideDownBinaryTree(root.Left)
	// 深度优先的方式进行寻找原root最左子树的左节点
	// 剩余步骤可以参考上方注释
	root.Left.Left = upsideDownBinaryTree(root.Right)
	root.Left.Right = root
	root.Left = nil
	root.Right = nil
	return newRoot
}

// @lc code=end
