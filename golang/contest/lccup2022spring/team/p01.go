package team

func getMinimumTime(time []int, fruits [][]int, limit int) int {
	ans := 0
	for _, fruit := range fruits {
		t := time[fruit[0]]
		num := fruit[1] / limit
		if fruit[1]%limit != 0 {
			num++
		}
		ans += num * t
	}
	return ans
}
