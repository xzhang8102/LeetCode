package golang

/*
 * @lc app=leetcode.cn id=1994 lang=golang
 *
 * [1994] The Number of Good Subsets
 */

// @lc code=start
func numberOfGoodSubsets(nums []int) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	const mod int = 1e9 + 7
	freq := [31]int{}
	for _, n := range nums {
		freq[n]++
	}
	const mask int = 1 << 10
	f := make([]int, mask)
	f[0] = 1
	for i := 2; i < 31; i++ {
		if freq[i] == 0 {
			continue
		}
		// prime factorization 质因数分解
		curr, x := 0, i
		valid := true
		for pos, prime := range primes {
			for x%prime == 0 {
				bit := (curr >> pos) & 1
				if bit == 1 {
					valid = false
					break
				}
				curr |= 1 << pos
				x /= prime
			}
			if !valid {
				break
			}
		}
		if !valid {
			continue
		}
		for prev := mask - 1; prev >= 0; prev-- {
			if (prev & curr) == 0 {
				f[prev|curr] = (f[prev|curr] + f[prev]*freq[i]) % mod
			}
		}
	}
	ans := 0
	for i := 1; i < mask; i++ {
		ans = (ans + f[i]) % mod
	}
	for i := 1; i <= freq[1]; i++ {
		ans = ans * 2 % mod
	}
	return ans
}

// @lc code=end
