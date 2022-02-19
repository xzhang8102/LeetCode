package golang

/*
 * @lc app=leetcode.cn id=969 lang=golang
 *
 * [969] Pancake Sorting
 */

// @lc code=start
func pancakeSort(arr []int) []int {
	ans := []int{}
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		maxIdx := 0
		for j := 0; j <= i; j++ {
			if arr[j] >= arr[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			lc969Reverse(arr[:maxIdx+1])
			lc969Reverse(arr[:i+1])
			ans = append(ans, maxIdx+1, i+1)
		}
	}
	return ans
}

func lc969Reverse(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

// @lc code=end
