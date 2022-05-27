package crack

// https://leetcode.cn/problems/find-closest-lcci/

func findClosest(words []string, word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}
	p1, p2 := -1, -1
	ans := len(words)
	for i, word := range words {
		if word == word1 {
			p1 = i
		}
		if word == word2 {
			p2 = i
		}
		if p1 != -1 && p2 != -1 {
			ans = min(abs(p1-p2), ans)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
