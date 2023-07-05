package main

import "testing"

func Test_kangaroo(t *testing.T) {
	type args struct {
		x1 int32
		v1 int32
		x2 int32
		v2 int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				x1: 0,
				v1: 3,
				x2: 4,
				v2: 2,
			},
			want: "YES",
		},
		{
			name: "Test case 2",
			args: args{
				x1: 0,
				v1: 2,
				x2: 5,
				v2: 3,
			},
			want: "NO",
		},
		{
			name: "Test case 3",
			args: args{
				x1: 43,
				v1: 2,
				x2: 70,
				v2: 2,
			},
			want: "NO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kangaroo(tt.args.x1, tt.args.v1, tt.args.x2, tt.args.v2); got != tt.want {
				t.Errorf("kangaroo() = %v, want %v", got, tt.want)
			}
		})
	}
}
