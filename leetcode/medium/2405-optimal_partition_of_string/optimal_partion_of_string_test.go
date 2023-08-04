package optimalpartitionofstring

import "testing"

func Test_partitionString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				s: "abacaba",
			},
			want: 4,
		},
		{
			name: "Test Case 2",
			args: args{
				s: "ssssss",
			},
			want: 6,
		},
		{
			name: "Test Case 3",
			args: args{
				s: "hdklqkcssgxlvehva",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionString(tt.args.s); got != tt.want {
				t.Errorf("partitionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
