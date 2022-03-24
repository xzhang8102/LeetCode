package golang

/*
 * @lc app=leetcode.cn id=818 lang=golang
 *
 * [818] Race Car
 */

// @lc code=start
func racecar(target int) int {
	type record struct {
		speed, location int
		op              byte
	}
	ans := -1
	if target == 0 {
		return 0
	}
	q := []record{{1, 0, 'A'}, {1, 0, 'R'}}
	for len(q) > 0 {
		tmp := q
		q = []record{}
		ans++
		for _, r := range tmp {
			if r.location == target {
				return ans
			}
			if r.op == 'A' {
				q = append(q, record{r.speed * 2, r.location + r.speed, 'A'})
				q = append(q, record{r.speed * 2, r.location + r.speed, 'R'})
			} else {
				if r.speed > 0 {
					q = append(q, record{-1, r.location, 'A'})
					q = append(q, record{-1, r.location, 'R'})
				} else {
					q = append(q, record{1, r.location, 'A'})
					q = append(q, record{1, r.location, 'R'})
				}
			}
		}
	}
	return ans
}

func lc818Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
