package golang

/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] Range Sum Query - Mutable
 */

// @lc code=start
type NumArray struct {
	data   []int
	presum []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	presum := make([]int, n)
	presum[0] = nums[0]
	for i := 1; i < n; i++ {
		presum[i] = presum[i-1] + nums[i]
	}
	return NumArray{
		nums,
		presum,
	}
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.data[index]
	for i := index; i < len(this.data); i++ {
		this.presum[i] += diff
	}
	this.data[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.presum[right]
	} else {
		return this.presum[right] - this.presum[left-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// @lc code=end
