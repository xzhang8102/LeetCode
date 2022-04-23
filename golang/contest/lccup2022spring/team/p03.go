package team

import "sort"

func getMaximumNumber(moles [][]int) int {
	a := [][]int{}
	flag := false
	for _, mole := range moles {
		if mole[0] == 0 {
			if mole[1] == 1 && mole[2] == 1 {
				flag = true
			}
		} else {
			a = append(a, mole)
		}
	}
	a = append(a, []int{0, 1, 1})
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	n := len(a)
	dp := make([]int, n)
	prefixMax := make([]int, n)
	ans := 0
	for i := 1; i < n; i++ {
		dp[i] = -1e8
		for j := i - 1; j >= 0; j-- {
			t := a[i][0] - a[j][0]
			d := abs(a[i][1]-a[j][1]) + abs(a[i][2]-a[j][2])
			if t > 4 {
				dp[i] = max(dp[i], prefixMax[j]+1)
				break
			}
			if d <= t {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
		prefixMax[i] = max(prefixMax[i-1], dp[i])
	}
	if flag {
		ans++
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
