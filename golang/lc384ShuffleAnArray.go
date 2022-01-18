package golang

import (
	"math/rand"
	"sort"
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

func Constructor(nums []int) Solution {
	return Solution{
		nums,
	}
}

func (this *Solution) Reset() []int {
	return this.data
}

func (this *Solution) Shuffle() []int {
	cache := map[float64]int{}
	for _, v := range this.data {
		cache[rand.Float64()] = v
	}
	keys := []float64{}
	for k := range cache {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	res := []int{}
	for _, k := range keys {
		res = append(res, cache[k])
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
