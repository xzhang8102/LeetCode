package weekly291

import (
	"fmt"
)

func removeDigit(number string, digit byte) string {
	ans := ""
	n := len(number)
	for i := 0; i < n; i++ {
		if number[i] == digit {
			tmp := fmt.Sprintf("%s%s", number[:i], number[i+1:])
			if ans == "" {
				ans = tmp
			} else if ans != tmp {
				for j := 0; j < n-1; j++ {
					if tmp[j] > ans[j] {
						ans = tmp
						break
					} else if tmp[j] < ans[j] {
						break
					}
				}
			}
		}
	}
	return ans
}
