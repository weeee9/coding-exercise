package main

import "testing"

func Test_timeConversion(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				s: "07:05:45PM",
			},
			want: "19:05:45",
		},
		{
			name: "Test Case 2",
			args: args{
				s: "12:40:22AM",
			},
			want: "00:40:22",
		},
		{
			name: "Test Case 3",
			args: args{
				s: "12:45:54PM",
			},
			want: "12:45:54",
		},
		{
			name: "Test Case 4",
			args: args{
				s: "06:40:00AM",
			},
			want: "06:40:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeConversion(tt.args.s); got != tt.want {
				t.Errorf("timeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
