package golang

import "math/rand"

/*
 * @lc app=leetcode.cn id=398 lang=golang
 *
 * [398] Random Pick Index
 */

// @lc code=start
type Solution []int

func lc398Constructor(nums []int) Solution {
	return nums
}

func (this *Solution) Pick(target int) int {
	cnt := 0
	ans := 0
	for i, num := range *this {
		if num == target {
			cnt++
			if rand.Intn(cnt) == 0 {
				ans = i
			}
		}
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end
