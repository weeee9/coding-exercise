package peakindexinamountainarray

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				A: []int{0, 1, 0},
			},
			want: 1,
		},
		{
			name: "Test Case 2",
			args: args{
				A: []int{0, 2, 1, 0},
			},
			want: 1,
		},
		{
			name: "Test Case 3",
			args: args{
				A: []int{0, 10, 5, 2},
			},
			want: 1,
		},
		{
			name: "Test Case 4",
			args: args{
				A: []int{3, 4, 5, 1},
			},
			want: 2,
		},
		{
			name: "Test Case 5",
			args: args{
				A: []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19},
			},
			want: 2,
		},
		{
			name: "Test Case 6",
			args: args{
				A: []int{1, 2, 3, 4, 5, 6, 7, 6, 5},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.A); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
