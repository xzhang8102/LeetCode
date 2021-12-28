package golang

import (
	"log"
	"testing"
)

func TestCanBeTypedWords(t *testing.T) {
	text := "leet code"
	brokenLetters := "lt"
	actual := canBeTypedWords(text, brokenLetters)
	expected := 1
	if actual == expected {
		log.Println("Success!")
	} else {
		log.Fatalf("Expected: %d, got: %d\n", expected, actual)
	}
}
