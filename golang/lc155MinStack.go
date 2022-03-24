package golang

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	stack []int
	min   int
}

func lc155Constructor() MinStack {
	return MinStack{
		[]int{},
		0,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.min = val
	}
	this.stack = append(this.stack, val)
	if val < this.min {
		this.stack = append(this.stack, this.min-val)
		this.min = val
	} else {
		this.stack = append(this.stack, 0)
	}
}

func (this *MinStack) Pop() {
	if diff := this.stack[len(this.stack)-1]; diff > 0 {
		this.min = this.stack[len(this.stack)-2] + diff
	}
	this.stack = this.stack[:len(this.stack)-2]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-2]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
