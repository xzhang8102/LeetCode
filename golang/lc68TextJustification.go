package golang

import "strings"

/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] Text Justification
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	ans := []string{}
	width := 0
	left, right := 0, 0
	for right < len(words) {
		if width == 0 || width > 0 && width+len(words[right])+right-left <= maxWidth {
			width += len(words[right])
			right++
		} else {
			tmp := ""
			totalGaps := maxWidth - width
			extraGapNum := 0
			gapVal := 0
			// more than one word in a line
			if right-left > 1 {
				extraGapNum = totalGaps % (right - left - 1)
				gapVal = totalGaps / (right - left - 1)
			}
			for i := left; i < right; i++ {
				if i > left {
					if extraGapNum > 0 {
						tmp += strings.Repeat(" ", gapVal+1)
						extraGapNum--
						totalGaps -= gapVal + 1
					} else {
						tmp += strings.Repeat(" ", gapVal)
						totalGaps -= gapVal
					}
				}
				tmp += words[i]
			}
			if totalGaps > 0 {
				tmp += strings.Repeat(" ", totalGaps)
			}
			ans = append(ans, tmp)
			left = right
			width = 0
		}
	}
	// if there are still words that are processed
	if width > 0 {
		tmp := strings.Join(words[left:], " ")
		if len(tmp) < maxWidth {
			tmp += strings.Repeat(" ", maxWidth-len(tmp))
		}
		ans = append(ans, tmp)
	}
	return ans
}

// @lc code=end
