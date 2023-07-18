package match

import "testing"

func Test_countMatches(t *testing.T) {
	type args struct {
		items     [][]string
		ruleKey   string
		ruleValue string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				items: [][]string{
					{"phone", "blue", "pixel"},
					{"computer", "silver", "lenovo"},
					{"phone", "gold", "iphone"},
				},
				ruleKey:   "color",
				ruleValue: "silver",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMatches(tt.args.items, tt.args.ruleKey, tt.args.ruleValue); got != tt.want {
				t.Errorf("countMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
