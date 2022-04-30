package golang

/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] Range Sum Query - Mutable
 */

// @lc code=start
type NumArray struct {
	nums []int
	tree []int
}

func Constructor(nums []int) NumArray {
	arr := NumArray{
		nums,
		make([]int, len(nums)+1),
	}
	for i, num := range nums {
		arr.treeUpdate(i+1, num)
	}
	return arr
}

func (this *NumArray) Update(index int, val int) {
	this.treeUpdate(index+1, val-this.nums[index])
	this.nums[index] = val
}

func (bit *NumArray) treeUpdate(index, diff int) {
	for index < len(bit.tree) {
		bit.tree[index] += diff
		index += index & (-index)
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.treeQuery(right+1) - this.treeQuery(left)
}

func (bit *NumArray) treeQuery(index int) int {
	ans := 0
	for index > 0 {
		ans += bit.tree[index]
		index -= index & (-index)
	}
	return ans
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// @lc code=end
