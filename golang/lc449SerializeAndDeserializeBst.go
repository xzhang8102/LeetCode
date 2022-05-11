package golang

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=449 lang=golang
 *
 * [449] Serialize and Deserialize BST
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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	data := []string{}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		data = append(data, strconv.Itoa(node.Val))
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return strings.Join(data, " ")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	s := strings.Split(data, " ")
	vals := make([]int, len(s))
	for i, str := range s {
		vals[i], _ = strconv.Atoi(str)
	}
	var build func(start, end int) *TreeNode
	build = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		root := &TreeNode{Val: vals[start]}
		next := start + 1
		for ; next <= end; next++ {
			if vals[next] > vals[start] {
				break
			}
		}
		root.Left = build(start+1, next-1)
		root.Right = build(next, end)
		return root
	}
	return build(0, len(vals)-1)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
// @lc code=end
