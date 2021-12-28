package golang

import "testing"

func TestLC3(t *testing.T) {
	want := 1
	if got := lengthOfLongestSubstring("bbbbb"); got != want {
		t.Errorf("Got: %d, want: %d", got, want)
	}
}
