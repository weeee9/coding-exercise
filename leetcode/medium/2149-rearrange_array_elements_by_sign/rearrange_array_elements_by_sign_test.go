package rearrangearrayelementsbysign

import (
	"reflect"
	"testing"
)

func Test_rearrangeArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{3, 1, -2, -5, 2, -4},
			},
			want: []int{3, -2, 1, -5, 2, -4},
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{1, 2, 3, -4, -1, -4},
			},
			want: []int{1, -4, 2, -1, 3, -4},
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{-1, 1},
			},
			want: []int{1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

