package golang

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	minStack []int
	data     []int
}

func Constructor() MinStack {
	return MinStack{
		[]int{},
		[]int{},
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	val := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if min := this.minStack[len(this.minStack)-1]; val == min {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
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
