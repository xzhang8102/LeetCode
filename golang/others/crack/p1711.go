package crack

// https://leetcode.cn/problems/find-closest-lcci/

func findClosest(words []string, word1 string, word2 string) int {
	dict := map[string][]int{}
	for i, word := range words {
		dict[word] = append(dict[word], i)
	}
	ans := len(words)
	i, j := 0, 0
	s1, s2 := dict[word1], dict[word2]
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			return 0
		}
		if s1[i] < s2[j] {
			ans = min(ans, s2[j]-s1[i])
			i++
		} else {
			ans = min(ans, s1[i]-s2[j])
			j++
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
