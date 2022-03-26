package golang

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
	in  []int
	out []int
}

func Constructor() MyStack {
	return MyStack{[]int{}, []int{}}
}

func (this *MyStack) Push(x int) {
	for _, v := range this.in {
		this.out = append(this.out, v)
	}
	this.in = []int{x}
}

func (this *MyStack) Pop() int {
	top := this.Top()
	this.in = []int{}
	for len(this.out) > 1 {
		this.in = append(this.in, this.out[0])
		this.out = this.out[1:]
	}
	this.in, this.out = this.out, this.in
	return top
}

func (this *MyStack) Top() int {
	return this.in[0]
}

func (this *MyStack) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
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
