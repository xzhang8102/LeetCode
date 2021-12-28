package golang

import "testing"

func TestLC212(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	findWords(board, words)
}
