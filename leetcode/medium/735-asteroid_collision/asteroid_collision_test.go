package asteroidcollision

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				asteroids: []int{5, 10, -5},
			},
			want: []int{5, 10},
		},
		{
			name: "Test Case 2",
			args: args{
				asteroids: []int{8, -8},
			},
			want: []int{},
		},
		{
			name: "Test Case 3",
			args: args{
				asteroids: []int{10, 2, -5},
			},
			want: []int{10},
		},
		{
			name: "Test Case 4",
			args: args{
				asteroids: []int{1, 2, 3, 4, -5},
			},
			want: []int{-5},
		},
		{
			name: "Test Case 5",
			args: args{
				asteroids: []int{-5, 10},
			},
			want: []int{-5, 10},
		},
		{
			name: "Test Case 6",
			args: args{
				asteroids: []int{2, 1, -1, -2},
			},
			want: []int{},
		},
		{
			name: "Test Case 7",
			args: args{
				asteroids: []int{-2, -1, 1, 2},
			},
			want: []int{-2, -1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_collide(t *testing.T) {
	type args struct {
		asteroids []int
		current   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				asteroids: []int{-1, -2, -3, -4},
				current:   5,
			},
			want: []int{-1, -2, -3, -4, 5},
		},
		{
			name: "Test Case 2",
			args: args{
				asteroids: []int{1, 2, 3, 4, 5},
				current:   -8,
			},
			want: []int{-8},
		},
		{
			name: "Test Case 3",
			args: args{
				asteroids: []int{1, 2, 3, 4, 5},
				current:   8,
			},
			want: []int{1, 2, 3, 4, 5, 8},
		},
		{
			name: "Test Case 4",
			args: args{
				asteroids: []int{5, 10},
				current:   -5,
			},
			want: []int{5, 10},
		},
		{
			name: "Test Case 5",
			args: args{
				asteroids: []int{5, 8, 6},
				current:   -7,
			},
			want: []int{5, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCollide(tt.args.asteroids, tt.args.current); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collide() = %v, want %v", got, tt.want)
			}
		})
	}
}
