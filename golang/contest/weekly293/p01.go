package weekly293

func removeAnagrams(words []string) []string {
	n := len(words)
	ans := []string{}
	var prev [26]int
	for i := 0; i < n; {
		cnt := [26]int{}
		if i == 0 {
			for _, ch := range words[i] {
				cnt[ch-'a']++
			}
			prev = cnt
			ans = append(ans, words[i])
			i++
			continue
		} else {
			j := i
			for j < n {
				for _, ch := range words[j] {
					cnt[ch-'a']++
				}
				if cnt != prev {
					prev = cnt
					ans = append(ans, words[j])
					break
				} else {
					cnt = [26]int{}
					j++
				}
			}
			i = j
		}
	}
	return ans
}
