package golang

func diStringMatch(s string) []int {
	n := len(s)
	ans := make([]int, n+1)
	lo, hi := 0, n
	for i, ch := range s {
		if ch == 'I' {
			ans[i] = lo
			lo++
		} else {
			ans[i] = hi
			hi--
		}
	}
	ans[n] = lo
	return ans
}
