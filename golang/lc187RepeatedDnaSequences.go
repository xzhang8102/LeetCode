package golang

/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	if n < 10 {
		return nil
	}
	mapping := map[byte]int{'A': 0x00, 'C': 0x01, 'G': 0x02, 'T': 0x03}
	cnt := map[int]int{}
	num := 0
	ans := []string{}
	for left, right := 0, 0; right <= n; right++ {
		if right-left < 10 {
			num = num<<2 | mapping[s[right]]
		} else {
			cnt[num]++
			if cnt[num] == 2 {
				ans = append(ans, s[left:right])
			}
			if right < n {
				num = (num<<2 | mapping[s[right]]) & 0xFFFFF
			}
			left++
		}
	}
	return ans
}

// @lc code=end
