package golang

/*
 * @lc app=leetcode.cn id=1220 lang=golang
 *
 * [1220] Count Vowels Permutation
 */

// @lc code=start
const mod int = 1e9 + 7

type lc1220Matrix [5][5]int

func (a lc1220Matrix) mul(b lc1220Matrix) lc1220Matrix {
	res := lc1220Matrix{}
	for i, row := range a {
		for j := range b[0] {
			for k, v := range row {
				res[i][j] = (res[i][j] + v*b[k][j]) % mod
			}
		}
	}
	return res
}

func (a lc1220Matrix) pow(n int) lc1220Matrix {
	res := lc1220Matrix{}
	for i := range res {
		res[i][i] = 1
	}
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res.mul(a)
		}
		if n > 1 {
			a = a.mul(a)
		}
	}
	return res
}

func countVowelPermutation(n int) int {
	m := lc1220Matrix{
		{0, 1, 0, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0},
	}
	mut := m.pow(n - 1)
	ans := 0
	for _, row := range mut {
		for _, v := range row {
			ans = (ans + v) % mod
		}
	}
	return ans
}

// @lc code=end
