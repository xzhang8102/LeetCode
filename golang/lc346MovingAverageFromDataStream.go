package golang

import "container/list"

/*
 * @lc app=leetcode.cn id=346 lang=golang
 *
 * [346] Moving Average from Data Stream
 */

// @lc code=start
type MovingAverage struct {
	window, sum int
	data        *list.List
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size,
		0,
		list.New(),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.data.Len() < this.window {
		this.sum += val
	} else {
		front := this.data.Front()
		diff := val - front.Value.(int)
		this.data.Remove(front)
		this.sum += diff
	}
	this.data.PushBack(val)
	return float64(this.sum) / float64(this.data.Len())
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
// @lc code=end
