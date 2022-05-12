package crack

// https://leetcode.cn/problems/one-away-lcci/

func oneEditAway(first string, second string) bool {
	n, m := len(first), len(second)
	if n < m {
		return oneEditAway(second, first)
	}
	if n-m > 1 {
		return false
	}
	for i, j := 0, 0; i < n && j < m; i, j = i+1, j+1 {
		if first[i] != second[j] {
			if n == m {
				return first[i+1:] == second[j+1:]
			} else {
				return first[i+1:] == second[j:]
			}
		}
	}
	return true
}
