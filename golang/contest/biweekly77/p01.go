package biweekly77

func countPrefixes(words []string, s string) int {
	dict := map[string]bool{}
	for i := 1; i <= len(s); i++ {
		dict[s[:i]] = true
	}
	ans := 0
	for _, word := range words {
		if dict[word] {
			ans++
		}
	}
	return ans
}
