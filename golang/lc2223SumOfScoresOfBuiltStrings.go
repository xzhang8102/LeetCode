package golang

/*
 * @lc app=leetcode.cn id=2223 lang=golang
 *
 * [2223] Sum of Scores of Built Strings
 */

// @lc code=start
func sumScores(s string) int64 {
	mod, base := int(1e9+7), 31
	n := len(s)
	prefix, mul := make([]int, n+1), make([]int, n+1)
	mul[0] = 1
	for i := 1; i <= n; i++ {
		prefix[i] = (prefix[i-1]*base + int(s[i-1]-'a'+1)) % mod
		mul[i] = mul[i-1] * base % mod
	}
	ans := int64(n)
	// m: suffix length
	for m := n - 1; m > 0; m-- {
		if s[n-m] != s[0] {
			continue
		}
		lo, hi := 1, m+1
		for lo < hi {
			mid := lo + (hi-lo)>>1
			hash := (prefix[n-m+mid] + mod - prefix[n-m]*mul[mid]%mod) % mod
			if hash == prefix[mid] {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		ans += int64(lo - 1)
	}
	return ans
}

// @lc code=end
