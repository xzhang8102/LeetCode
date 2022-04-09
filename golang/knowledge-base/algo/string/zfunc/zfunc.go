package zfunc

// ref:
// https://oi-wiki.org/string/z-func/
// https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/strings.go

// func calcZ(s string) []int {
// 	n := len(s)
// 	z := make([]int, n)
// 	for i, l, r := 1, 0, 0; i < n; i++ {
// 		if i <= r {
// 			z[i] = min(z[i-l], r-i+1)
// 		}
// 		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
// 			l, r = i, i+z[i]
// 			z[i]++
// 		}
// 	}
// 	z[0] = n
// 	return z
// }

// calculate z-values
// z-value: the length of the longest prefix of
// the suffix str[i:] which is also the prefix of the str
func calcZ(str string) []int {
	n := len(str)
	z := make([]int, n)
	left, right := 0, 0
	for i := 1; i < n; i++ {
		if i > right {
			left, right = i, i
			for right < n && str[right] == str[right-left] {
				right++
			}
			z[i] = right - left
			right--
		} else {
			if i+z[i-left]-1 < right {
				z[i] = z[i-left]
			} else {
				left = i
				for right < n && str[right] == str[right-left] {
					right++
				}
				z[i] = right - left
				right--
			}
		}
	}
	return z
}

func ZSearch(str, pattern string) []int {
	// separate pattern and str with a special character
	z := calcZ(pattern + "$" + str)
	n := len(pattern)
	ans := []int{}
	for i, v := range z[n+1:] {
		if v == n {
			ans = append(ans, i)
		}
	}
	return ans
}
