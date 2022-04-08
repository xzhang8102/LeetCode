package kmp

import (
	"reflect"
	"testing"
)

func Test_build(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "build prefix array for KMP",
			args: args{
				pattern: "acacabacacabacacac",
			},
			want: []int{0, 0, 1, 2, 3, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := build(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMP(t *testing.T) {
	type args struct {
		str     string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{
				str:     "ABC ABCDAB ABCDABCDABDE",
				pattern: "ABCDABD",
			},
			want: []int{15},
		},
		{
			name: "test case 2",
			args: args{
				"AABAACAADAABAABA",
				"AABA",
			},
			want: []int{0, 9, 12},
		},
		{
			name: "edge case",
			args: args{
				"ab",
				"",
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KMP(tt.args.str, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KMP() = %v, want %v", got, tt.want)
			}
		})
	}
}
