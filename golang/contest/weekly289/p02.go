package weekly289

func minimumRounds(tasks []int) int {
	ans := 0
	cache := map[int]int{1: -1, 2: 1, 3: 1}
	var findWays func(n int) int
	findWays = func(n int) int {
		if cache[n] != 0 {
			return cache[n]
		}
		for i := n / 3; i >= 0; i-- {
			remain := n - i*3
			if remain%2 == 0 {
				cache[n] = i + remain/2
				break
			}
		}
		return cache[n]
	}
	dict := map[int]int{}
	for _, task := range tasks {
		dict[task]++
	}
	for _, n := range dict {
		ways := findWays(n)
		if ways == -1 {
			return -1
		}
		ans += ways
	}
	return ans
}
