package golang

/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */

// @lc code=start
type NumArray struct {
	data []int
	sum  []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	arr := make([]int, n)
	arr[0] = nums[0]
	for i := 1; i < n; i++ {
		arr[i] = arr[i-1] + nums[i]
	}
	return NumArray{
		nums,
		arr,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sum[right]
	}
	return this.sum[right] - this.sum[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end
