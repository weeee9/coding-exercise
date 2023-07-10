package sorting

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "quickSort",
			args: args{
				array: []int{5, 2, 4, 6, 1, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "quickSort",
			args: args{
				array: []int{5, 2},
			},
			want: []int{2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
