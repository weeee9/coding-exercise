package main

import (
	"testing"
)

func Test_miniMaxSum(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				arr: []int32{1, 2, 3, 4, 5},
			},
		},
		{
			name: "Test case 2",
			args: args{
				arr: []int32{7, 69, 2, 221, 8974},
			},
		},
		{
			name: "Test case 3",
			args: args{
				arr: []int32{793810624, 895642170, 685903712, 623789054, 468592370},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			miniMaxSum(tt.args.arr)
		})
	}
}
