package pairSum

import "testing"

func Test_pairSum(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairSum(tt.args.head); got != tt.want {
				t.Errorf("pairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
