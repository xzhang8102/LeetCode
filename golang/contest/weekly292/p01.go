package weekly292

func largestGoodInteger(num string) string {
	n := len(num)
	max := ""
	for i := 0; i < n-2; i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] {
			if max == "" || max[0] < num[i] {
				max = num[i : i+3]
			}
		}
	}
	return max
}
