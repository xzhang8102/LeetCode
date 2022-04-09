package zfunc

import (
	"reflect"
	"testing"
)

func TestZSearch(t *testing.T) {
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
			name: "test case",
			args: args{
				str:     "xabcabzabc",
				pattern: "abc",
			},
			want: []int{1, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZSearch(tt.args.str, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcZ(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{
				"aabxaabxcaabxaabxay",
			},
			want: []int{0, 1, 0, 0, 4, 1, 0, 0, 0, 8, 1, 0, 0, 5, 1, 0, 0, 1, 0},
		},
		{
			name: "test case 2",
			args: args{
				"abaxabab",
			},
			want: []int{0, 0, 1, 0, 3, 0, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcZ(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcZ() = %v, want %v", got, tt.want)
			}
		})
	}
}
