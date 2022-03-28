package golang

/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] Scramble String
 */

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	const (
		N       = -1
		Y       = 1
		PENDING = 0
	)
	n := len(s1)
	cache := make([][][]int, n)
	for i := range cache {
		cache[i] = make([][]int, n)
		for j := range cache[i] {
			cache[i][j] = make([]int, n+1)
		}
	}
	var dfs func(i, j, length int) bool
	dfs = func(i, j, length int) bool {
		if cache[i][j][length] != PENDING {
			return cache[i][j][length] == Y
		}
		sub1, sub2 := s1[i:i+length], s2[j:j+length]
		if sub1 == sub2 {
			cache[i][j][length] = Y
			return true
		}
		if !lc87Check(sub1, sub2) {
			cache[i][j][length] = N
			return false
		}
		for k := 1; k < length; k++ {
			if dfs(i, j, k) && dfs(i+k, j+k, length-k) {
				cache[i][j][length] = Y
				return true
			}
			if dfs(i, length-k+j, k) && dfs(i+k, j, length-k) {
				cache[i][j][length] = Y
				return true
			}
		}
		cache[i][j][length] = N
		return false
	}
	return dfs(0, 0, n)
}

func lc87Check(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	dict := [26]int{}
	for i := range s1 {
		dict[s1[i]-'a']++
		dict[s2[i]-'a']--
	}
	for _, v := range dict {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
