package search

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "Test Case 2",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "Test Case 3",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "Test Case 4",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "Test Case 4",
			args: args{
				nums:   []int{3, 1},
				target: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPivot(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 4,
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: -1,
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{1},
			},
			want: -1,
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{1, 2},
			},
			want: -1,
		},
		{
			name: "Test Case 5",
			args: args{
				nums: []int{7, 1, 2, 3, 4, 5, 6},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPivot(tt.args.nums); got != tt.want {
				t.Errorf("findPivot() = %v, want %v", got, tt.want)
			}
		})
	}
}
