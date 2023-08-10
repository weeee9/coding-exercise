package reversesubstringsbetweeneachpairofparentheses

import (
	"testing"
)

func Test_reverseParentheses(t *testing.T) {
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
				s: "(abcd)",
			},
			want: "dcba",
		},
		{
			name: "Test Case 2",
			args: args{
				s: "(u(love)i)",
			},
			want: "iloveu",
		},
		{
			name: "Test Case 3",
			args: args{
				s: "(ed(et(oc))el)",
			},
			want: "leetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseParentheses(tt.args.s); got != tt.want {
				t.Errorf("reverseParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		s     []rune
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				s:     []rune("acbd"),
				left:  1,
				right: 2,
			},
			want: "abcd",
		},
		{
			name: "Test Case 2",
			args: args{
				s:     []rune("leetedoc"),
				left:  4,
				right: 7,
			},
			want: "leetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
