package golang

/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start
func reverseVowels(s string) string {
	n := len(s)
	buf := []byte(s)
	for left, right := 0, n-1; left < right; {
		for left < right && !lc345CheckVowel(buf[left]) {
			left++
		}
		if left == right {
			break
		}
		for left < right && !lc345CheckVowel(buf[right]) {
			right--
		}
		buf[left], buf[right] = buf[right], buf[left]
		left++
		right--
	}
	return string(buf)
}

func lc345CheckVowel(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

// @lc code=end
