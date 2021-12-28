package golang

import "testing"

func TestLC7(t *testing.T) {
	want := -2143847412
	if got := reverse(-2147483412); got != want {
		t.Errorf("Got: %d, Not %d", got, want)
	}
}
