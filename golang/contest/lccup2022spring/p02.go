package lccup2022spring

func perfectMenu(materials []int, cookbooks [][]int, attribute [][]int, limit int) int {
	ans := -1
	n := len(cookbooks)
	fullness, tasty := 0, 0
	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos == n {
			if fullness >= limit {
				if tasty > ans {
					ans = tasty
				}
			}
			return
		}
		for i := pos; i < n; i++ {
			valid := true
			tmp := materials
			materials = make([]int, len(materials))
			for j, cost := range cookbooks[i] {
				if tmp[j] < cost {
					valid = false
					break
				} else {
					materials[j] = tmp[j] - cost
				}
			}
			if !valid {
				materials = tmp
				continue
			}
			tasty += attribute[i][0]
			fullness += attribute[i][1]
			backtrack(i + 1)
			tasty -= attribute[i][0]
			fullness -= attribute[i][1]
			materials = tmp
		}
		if fullness >= limit {
			if tasty > ans {
				ans = tasty
			}
		}
	}
	backtrack(0)
	return ans
}
