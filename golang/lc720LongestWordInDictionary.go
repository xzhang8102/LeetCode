package golang

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] Longest Word in Dictionary
 */

// @lc code=start
func longestWord(words []string) string {
	dict := map[int]map[string]bool{}
	maxLen := 0
	for _, word := range words {
		length := len(word)
		if length > maxLen {
			maxLen = length
		}
		if dict[length] == nil {
			dict[length] = map[string]bool{word: true}
		} else {
			dict[length][word] = true
		}
	}
	candidates := dict[1]
	if candidates == nil {
		return ""
	}
	for i := 2; i <= maxLen; i++ {
		tmp := map[string]bool{}
		for s := range dict[i] {
			if candidates[s[:i-1]] {
				tmp[s] = true
			}
		}
		if len(tmp) == 0 {
			break
		} else {
			candidates = tmp
		}
	}
	ans := ""
	for s := range candidates {
		if ans == "" || s < ans {
			ans = s
		}
	}
	return ans
}

// @lc code=end
