package golang

import (
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] Shuffle an Array
 */

// @lc code=start
type Solution struct {
	data []int
}

func lc384Constructor(nums []int) Solution {
	return Solution{
		nums,
	}
}

func (this *Solution) Reset() []int {
	return this.data
}

func (this *Solution) Shuffle() []int {
	res := append([]int(nil), this.data...)
	for i := range res {
		j := i + rand.Intn(len(res)-i)
		res[i], res[j] = res[j], res[i]
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end
