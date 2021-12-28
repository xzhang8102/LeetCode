package golang

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
/*
[1,2,3,4,5,6,7]
k = 3
[#,#,#,1,2,3,4,5,6,7]
[5,6,7,1,2,3,4]
*/
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverseArr(nums)
	reverseArr(nums[:k])
	reverseArr(nums[k:])
}

func reverseArr(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// @lc code=end
