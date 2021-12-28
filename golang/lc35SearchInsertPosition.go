package golang

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	low, hi := 0, len(nums)-1
	for low <= hi {
		mid := (low + hi) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low > hi {
		return low
	} else {
		return hi
	}
}

// @lc code=end
