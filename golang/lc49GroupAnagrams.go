package golang

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	dict := map[[26]int][]string{}
	for _, str := range strs {
		pattern := [26]int{}
		for _, char := range str {
			pattern[char-'a']++
		}
		dict[pattern] = append(dict[pattern], str)
	}
	ans := make([][]string, 0, len(dict))
	for _, v := range dict {
		ans = append(ans, v)
	}
	return ans
}

// @lc code=end
