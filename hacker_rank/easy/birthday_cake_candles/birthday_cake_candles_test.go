package main

import "testing"

func Test_birthdayCakeCandles(t *testing.T) {
	type args struct {
		candles []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test case 1",
			args: args{
				candles: []int32{3, 2, 1, 3},
			},
			want: 2,
		},
		{
			name: "Test case 2",
			args: args{
				candles: []int32{4, 4, 1, 3},
			},
			want: 2,
		},
		{
			name: "Test case 3",
			args: args{
				candles: []int32{4, 1, 3},
			},
			want: 1,
		},
		{
			name: "Test case 4",
			args: args{
				candles: []int32{82, 49, 82, 82, 41, 82, 15, 63, 38, 25},
			},
			want: 4,
		},
		{
			name: "Test case 5",
			args: args{
				candles: []int32{44, 53, 31, 27, 77, 60, 66, 77, 26, 36},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := birthdayCakeCandles(tt.args.candles); got != tt.want {
				t.Errorf("birthdayCakeCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}
