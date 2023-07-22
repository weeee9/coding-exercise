package buildarray

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	type args struct {
		target []int
		n      int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test Case 1",
			args: args{
				target: []int{1, 3},
				n:      3,
			},
			want: []string{"Push", "Push", "Pop", "Push"},
		},
		{
			name: "Test Case 2",
			args: args{
				target: []int{1, 2, 3},
				n:      3,
			},
			want: []string{"Push", "Push", "Push"},
		},
		{
			name: "Test Case 3",
			args: args{
				target: []int{1, 2},
				n:      4,
			},
			want: []string{"Push", "Push"},
		},
		{
			name: "Test Case 4",
			args: args{
				target: []int{2, 3, 4},
				n:      4,
			},
			want: []string{"Push", "Pop", "Push", "Push", "Push"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray(tt.args.target, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
