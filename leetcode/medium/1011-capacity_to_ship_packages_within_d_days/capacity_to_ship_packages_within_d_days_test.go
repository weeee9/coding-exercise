package capacitytoshippackageswithinddays

import (
	"testing"
)

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		days    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				days:    5,
			},
			want: 15,
		},
		{
			name: "Test Case 2",
			args: args{
				weights: []int{3, 2, 2, 4, 1, 4},
				days:    3,
			},
			want: 6,
		},
		{
			name: "Test Case 3",
			args: args{
				weights: []int{1, 2, 3, 1, 1},
				days:    4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.days); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_possible(t *testing.T) {
	type args struct {
		weights []int
		cap     int
		days    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				cap:     32, // 10 + (55-10)/2
				days:    5,
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				cap:     21, // 10 + (32-10)/2
				days:    5,
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				cap:     15, // 10 + (21-10)/2
				days:    5,
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				cap:     12, // 10 + (15-10)/2
				days:    5,
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				cap:     13, // 12 + (15-12)/2
				days:    5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possible(tt.args.weights, tt.args.cap, tt.args.days); got != tt.want {
				t.Errorf("possible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMaxAndTotalLoad(t *testing.T) {
	type args struct {
		weights []int
	}
	tests := []struct {
		name      string
		args      args
		wantMax   int
		wantTotal int
	}{
		{
			name: "Test Case 1",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			wantMax:   10,
			wantTotal: 55,
		},
		{
			name: "Test Case 2",
			args: args{
				weights: []int{3, 2, 2, 4, 1, 4},
			},
			wantMax:   4,
			wantTotal: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotTotal := getMaxAndTotalLoad(tt.args.weights)
			if gotMax != tt.wantMax {
				t.Errorf("getMaxAndTotalLoad() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("getMaxAndTotalLoad() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
