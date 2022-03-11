package golang

/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] Corporate Flight Bookings
 */

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	for _, booking := range bookings {
		left, right, val := booking[0]-1, booking[1]-1, booking[2]
		diff[left] += val
		diff[right+1] -= val
	}
	ans := make([]int, n)
	ans[0] = diff[0]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + diff[i]
	}
	return ans
}

// @lc code=end
