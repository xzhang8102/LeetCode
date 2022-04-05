package golang

/*
 * @lc app=leetcode.cn id=346 lang=golang
 *
 * [346] Moving Average from Data Stream
 */

// @lc code=start
type MovingAverage struct {
	head, sum int
	data      []int
}

func lc346Constructor(size int) MovingAverage {
	return MovingAverage{
		0,
		0,
		make([]int, 0, size),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	size := len(this.data)
	if size < cap(this.data) {
		this.sum += val
		this.data = append(this.data, val)
	} else {
		diff := val - this.data[this.head]
		this.sum += diff
		this.head++
		if this.head == size {
			this.head = 0
		}
		tail := (this.head + size - 1) % size
		this.data[tail] = val
	}
	return float64(this.sum) / float64(len(this.data))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
// @lc code=end
