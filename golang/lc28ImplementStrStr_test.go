package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	h1 := "hello"
	n1 := "ll"
  assert.Equal(t, 2, strStr(h1, n1), "Index should be 2")
  h2 := "aaaaa"
  n2 := "aab"
  assert.Equal(t, -1, strStr(h2, n2), "Not found")
}
