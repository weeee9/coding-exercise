package convert

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findMatrix(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{1, 3, 4, 1, 2, 3, 1},
			},
			want: [][]int{{1, 2, 3, 4}, {1, 3}, {1}},
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{9, 8, 8, 4, 9, 8, 8, 3, 9},
			},
			want: [][]int{{3, 4, 8, 9}, {8, 9}, {8, 9}, {8}},
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{3, 5, 3, 3, 8, 3, 2, 5},
			},
			want: [][]int{{2, 3, 5, 8}, {3, 5}, {3}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMatrix(tt.args.nums)

			for i := range got {
				sort.Ints(got[i])
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anotherSolution(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{1, 3, 4, 1, 2, 3, 1},
			},
			want: [][]int{{1, 2, 3, 4}, {1, 3}, {1}},
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{9, 8, 8, 4, 9, 8, 8, 3, 9},
			},
			want: [][]int{{3, 4, 8, 9}, {8, 9}, {8, 9}, {8}},
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{3, 5, 3, 3, 8, 3, 2, 5},
			},
			want: [][]int{{2, 3, 5, 8}, {3, 5}, {3}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := anotherSolution(tt.args.nums)

			for i := range got {
				sort.Ints(got[i])
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anotherSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
