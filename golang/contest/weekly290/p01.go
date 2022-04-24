package weekly290

func intersection(nums [][]int) []int {
	freq := make([]int, 1001)
	n := len(nums)
	for _, num := range nums {
		for _, v := range num {
			freq[v]++
		}
	}
	ans := []int{}
	for i := 0; i <= 1000; i++ {
		if freq[i] == n {
			ans = append(ans, i)
		}
	}
	return ans
}
