package golang

/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] Word Ladder II
 */

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	set := map[string]bool{}
	for _, word := range wordList {
		set[word] = true
	}
	if len(set) == 0 || !set[endWord] {
		return nil
	}
	ans := [][]string{}
	graph := map[string][]string{}
	q := []struct {
		word string
		step int
	}{{beginWord, 1}}
	record := map[string]int{beginWord: 1}
	found := false
	for len(q) > 0 && !found {
		tmp := q
		q = nil
		for _, tuple := range tmp {
			word := tuple.word
			if word == endWord {
				found = true
				break
			}
			buf := []byte(word)
			for i := 0; i < len(buf); i++ {
				ch := buf[i]
				for b := byte('a'); b <= 'z'; b++ {
					buf[i] = b
					if next := string(buf); set[next] {
						if step, ok := record[next]; !ok {
							record[next] = tuple.step + 1
							q = append(q, struct {
								word string
								step int
							}{next, tuple.step + 1})
							graph[word] = append(graph[word], next)
						} else if step == tuple.step+1 {
							graph[word] = append(graph[word], next)
						}
					}
				}
				buf[i] = ch
			}
		}
	}
	pick := []string{beginWord}
	var dfs func(string)
	dfs = func(s string) {
		if s == endWord {
			tmp := make([]string, len(pick))
			copy(tmp, pick)
			ans = append(ans, tmp)
			return
		}
		for _, next := range graph[s] {
			pick = append(pick, next)
			dfs(next)
			pick = pick[:len(pick)-1]
		}
	}
	dfs(beginWord)
	return ans
}

// @lc code=end
