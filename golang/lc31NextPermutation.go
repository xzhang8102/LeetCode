package golang

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			next := i + 1
			// find the valid insert index (too hard if using binary search :/)
			for ; next < n; next++ {
				if nums[next] > nums[i] && (next == n-1 || nums[i] >= nums[next+1]) {
					break
				}
			}
			nums[i], nums[next] = nums[next], nums[i]
			lc31Reverse(nums[i+1:])
			break
		}
	}
	if i == -1 {
		lc31Reverse(nums)
	}
}

func lc31Reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

// 1 4 3 6 8 7 0 -1
//       |
// 1 4 3 7 8 6 0 -1
//       |
// 1 4 3 7 -1 0 6 8

// @lc code=end
