package golang

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=440 lang=golang
 *
 * [440] K-th Smallest in Lexicographical Order
 */

// @lc code=start
func lc440findKthNumber(n int, k int) int {
	ans := 1
	for k > 1 {
		cnt := lc440PrifixCnt(ans, n)
		if k > cnt {
			// accept the whole sub-trie
			ans++
			k -= cnt
		} else {
			// into the sub-trie
			ans *= 10
			// take original ans as a valid number
			k--
		}
	}
	return ans
}

// count all valid number (smaller than limit) with prefix
func lc440PrifixCnt(prefix, limit int) int {
	prefixStr, limitStr := strconv.Itoa(prefix), strconv.Itoa(limit)
	prefixLen, limitLen := len(prefixStr), len(limitStr)
	cnt := 0
	// add all the number has smaller value than limit
	for i := 0; i < limitLen-prefixLen; i++ {
		cnt += int(math.Pow10(i))
	}
	if x, _ := strconv.Atoi(limitStr[:prefixLen]); x > prefix {
		cnt += int(math.Pow10(limitLen - prefixLen))
	} else if x == prefix {
		cnt += limit - x*int(math.Pow10(limitLen-prefixLen)) + 1
	}
	return cnt
}

// @lc code=end
