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
