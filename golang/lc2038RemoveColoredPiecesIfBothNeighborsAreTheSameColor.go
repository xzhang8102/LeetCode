package golang

/*
 * @lc app=leetcode.cn id=2038 lang=golang
 *
 * [2038] Remove Colored Pieces if Both Neighbors are the Same Color
 */

// @lc code=start
func winnerOfGame(colors string) bool {
	freq := [2]int{}
	for left, right := 0, 0; right < len(colors); {
		if colors[right] == colors[left] {
			if right-left+1 >= 3 {
				freq[colors[right]-'A']++
			}
			right++
		} else {
			left = right
		}
	}
	return freq[0] > freq[1]
}

// @lc code=end
