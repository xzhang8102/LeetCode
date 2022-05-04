package golang

/*
 * @lc app=leetcode.cn id=1823 lang=golang
 *
 * [1823] Find the Winner of the Circular Game
 */

// @lc code=start
func findTheWinner(n int, k int) int {
	players := make([]int, n)
	for i := range players {
		players[i] = i + 1
	}
	next := 0
	for len(players) > 1 {
		size := len(players)
		next = (k + next - 1) % size
		copy(players[next:size-1], players[next+1:])
		players = players[:size-1]
	}
	return players[0]
}

// @lc code=end
