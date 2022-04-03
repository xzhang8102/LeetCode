package biweekly75

func numberOfWays(s string) int64 {
	n := len(s)
	if n == 0 {
		return 0
	}
	count := make([]int, n)
	for i, ch := range s {
		if ch == '1' {
			if i == 0 {
				count[i] = 1
			} else {
				count[i] = count[i-1] + 1
			}
		} else if i > 0 {
			count[i] = count[i-1]
		}
	}
	ans := int64(0)
	for i := 1; i < n-1; i++ {
		if s[i] == '1' {
			ans += int64(i-count[i-1]) * int64(n-1-i-(count[n-1]-count[i]))
		} else {
			ans += int64(count[i-1]) * int64(count[n-1]-count[i])
		}
	}
	return ans
}
