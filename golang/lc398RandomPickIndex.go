package golang

import "math/rand"

/*
 * @lc app=leetcode.cn id=398 lang=golang
 *
 * [398] Random Pick Index
 */

// @lc code=start
type Solution struct {
	data map[int][]int
}

func Constructor(nums []int) Solution {
	dict := map[int][]int{}
	for i, num := range nums {
		dict[num] = append(dict[num], i)
	}
	return Solution{dict}
}

func (this *Solution) Pick(target int) int {
	if freq := len(this.data[target]); freq == 1 {
		return this.data[target][0]
	} else {
		return this.data[target][rand.Intn(freq)]
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end
