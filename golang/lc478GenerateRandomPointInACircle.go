package golang

import (
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=478 lang=golang
 *
 * [478] Generate Random Point in a Circle
 */

// @lc code=start
type Solution struct {
	r, centerX, centerY float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius, x_center, y_center}
}

func (this *Solution) RandPoint() []float64 {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		if x*x+y*y < 1 {
			return []float64{this.centerX + x*this.r, this.centerY + y*this.r}
		}
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
// @lc code=end
