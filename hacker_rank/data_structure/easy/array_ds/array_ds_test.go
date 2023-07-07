package main

import (
	"reflect"
	"testing"
)

func Test_reverseArray(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "Test case 1",
			args: args{
				a: []int32{1, 4, 3, 2},
			},
			want: []int32{2, 3, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseArray(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
