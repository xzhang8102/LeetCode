package golang

func diStringMatch(s string) []int {
	n := len(s)
	picked := make([]bool, n+1)
	found := false
	var backtrack func(idx, prev int) []int
	backtrack = func(idx, prev int) []int {
		if idx == n {
			found = true
			return nil
		}
		var res []int
		if idx == -1 {
			for i := 0; i <= n; i++ {
				picked[i] = true
				tmp := backtrack(0, i)
				if found {
					res = append([]int{i}, tmp...)
					return res
				}
				picked[i] = false
			}
		} else {
			if s[idx] == 'I' {
				for i := prev + 1; i <= n; i++ {
					if !picked[i] {
						picked[i] = true
						tmp := backtrack(idx+1, i)
						if found {
							res = append([]int{i}, tmp...)
							return res
						}
						picked[i] = false
					}
				}
			} else {
				for i := prev - 1; i >= 0; i-- {
					if !picked[i] {
						picked[i] = true
						tmp := backtrack(idx+1, i)
						if found {
							res = append([]int{i}, tmp...)
							return res
						}
						picked[i] = false
					}
				}
			}
		}
		return res
	}
	return backtrack(-1, -1)
}
