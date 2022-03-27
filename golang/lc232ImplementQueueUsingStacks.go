package golang

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
	main, second []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.main = append(this.main, x)
}

func (this *MyQueue) Pop() int {
	if len(this.second) == 0 {
		for size := len(this.main); size > 0; size-- {
			this.second = append(this.second, this.main[size-1])
			this.main = this.main[:size-1]
		}
	}
	size := len(this.second)
	top := this.second[size-1]
	this.second = this.second[:size-1]
	return top
}

func (this *MyQueue) Peek() int {
	if len(this.second) == 0 {
		for size := len(this.main); size > 0; size-- {
			this.second = append(this.second, this.main[size-1])
			this.main = this.main[:size-1]
		}
	}
	size := len(this.second)
	top := this.second[size-1]
	return top
}

func (this *MyQueue) Empty() bool {
	return len(this.main) == 0 && len(this.second) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
