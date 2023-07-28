package queriesonnumberofpointsinsideacircle

import (
	"reflect"
	"testing"
)

func Test_countPoints(t *testing.T) {
	type args struct {
		points  [][]int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				points: [][]int{
					{1, 3},
					{3, 3},
					{5, 3},
					{2, 2},
				},
				queries: [][]int{
					{2, 3, 1},
					{4, 3, 1},
					{1, 1, 2},
				},
			},
			want: []int{3, 2, 2},
		},
		{
			name: "Test Case 2",
			args: args{
				points: [][]int{
					{1, 1},
					{2, 2},
					{3, 3},
					{4, 4},
					{5, 5},
				},
				queries: [][]int{
					{1, 2, 2},
					{2, 2, 2},
					{4, 3, 2},
					{4, 3, 3},
				},
			},
			want: []int{2, 3, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints(tt.args.points, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pointIsCircle(t *testing.T) {
	type args struct {
		circle []int
		point  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				circle: []int{2, 3, 1},
				point:  []int{1, 3},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				circle: []int{2, 3, 1},
				point:  []int{3, 3},
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				circle: []int{2, 3, 1},
				point:  []int{5, 3},
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				circle: []int{2, 3, 1},
				point:  []int{2, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointIsCircle(tt.args.circle, tt.args.point); got != tt.want {
				t.Errorf("pointIsCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
