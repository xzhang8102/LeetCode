package golang

import "math"

/*
 * @lc app=leetcode.cn id=1231 lang=golang
 *
 * [1231] Divide Chocolate
 */

// @lc code=start
func maximizeSweetness(sweetness []int, k int) int {
	// low bound and higher bound for sweetness
	min, sum := sweetness[0], sweetness[0]
	for _, s := range sweetness[1:] {
		if s < min {
			min = s
		}
		sum += s
	}
	ans := min
	for lo, hi := min, sum; lo <= hi; {
		mid := lo + (hi-lo)/2
		if mid*(k+1) > sum {
			hi = mid - 1
		} else {
			check := lc1232Cut(sweetness, mid, k+1)
			if check == -1 {
				hi = mid - 1
			} else {
				ans = check
				lo = check + 1
			}
		}
	}
	return ans
}

// manage to have the input value
// return the actual min value
func lc1232Cut(arr []int, wanted, k int) int {
	pieces := 0
	tmp := 0
	min := math.MaxInt32
	for _, val := range arr {
		tmp += val
		if tmp >= wanted {
			if tmp < min {
				min = tmp
			}
			pieces++
			tmp = 0
		}
	}
	if pieces >= k {
		return min
	} else {
		return -1
	}
}

// @lc code=end
