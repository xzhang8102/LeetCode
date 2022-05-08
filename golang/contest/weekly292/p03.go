package weekly292

func countTexts(pressedKeys string) int {
	mapping := map[byte]int{
		'2': 3, '3': 3, '4': 3, '5': 3, '6': 3, '7': 4, '8': 3, '9': 4,
	}
	n := len(pressedKeys)
	var mod int = 1e9 + 7
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]
		for j := i - 1; j > 0 && pressedKeys[j-1] == pressedKeys[i-1] && i-j+1 <= mapping[pressedKeys[i-1]]; j-- {
			dp[i] = (dp[i] + dp[j-1]) % mod
		}
	}
	return dp[n]
}
