package score

import (
	"reflect"
	"testing"
)

func Test_quickSortTheStudents(t *testing.T) {
	type args struct {
		score [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1",
			args: args{
				score: [][]int{
					{3, 91},
					{2, 97},
					{1, 99},
				},
				k: 1,
			},
			want: [][]int{
				{1, 99},
				{2, 97},
				{3, 91},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				score: [][]int{
					{10, 6, 9, 1},
					{7, 5, 11, 2},
					{4, 8, 3, 15},
				},
				k: 2,
			},
			want: [][]int{
				{7, 5, 11, 2},
				{10, 6, 9, 1},
				{4, 8, 3, 15},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSortTheStudents(tt.args.score, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSortTheStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertionSortTheStudents(t *testing.T) {
	type args struct {
		score [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1",
			args: args{
				score: [][]int{
					{3, 91},
					{2, 97},
					{1, 99},
				},
				k: 1,
			},
			want: [][]int{
				{1, 99},
				{2, 97},
				{3, 91},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				score: [][]int{
					{10, 6, 9, 1},
					{7, 5, 11, 2},
					{4, 8, 3, 15},
				},
				k: 2,
			},
			want: [][]int{
				{7, 5, 11, 2},
				{10, 6, 9, 1},
				{4, 8, 3, 15},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortTheStudents(tt.args.score, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortTheStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
