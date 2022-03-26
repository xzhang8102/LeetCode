package golang

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{[]int{}}
}

func (this *MyStack) Push(x int) {
	n := len(this.queue)
	this.queue = append(this.queue, x)
	for n > 0 {
		this.queue = append(this.queue, this.queue[0])
		this.queue = this.queue[1:]
		n--
	}
}

func (this *MyStack) Pop() int {
	top := this.Top()
	this.queue = this.queue[1:]
	return top
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
