package golang

import "math/rand"

/*
 * @lc app=leetcode.cn id=382 lang=golang
 *
 * [382] Linked List Random Node
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (this *Solution) GetRandom() int {
	ans := this.head.Val
	for ptr, idx := this.head, 1; ptr != nil; ptr, idx = ptr.Next, idx+1 {
		if rand.Intn(idx) == 0 {
			ans = ptr.Val
		}
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end
