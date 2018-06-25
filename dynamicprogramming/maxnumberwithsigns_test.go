package dynamicprogramming

import (
	"testing"
)

func TestSolveMaxNumberWithSignsRecursive(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1, 2",
			args: args{
				input: []int{1, 2},
			},
			want: 3,
		},
		{
			name: "1, 2, 3",
			args: args{
				input: []int{1, 2, 3},
			},
			want: 9,
		},
		{
			name: "1, 2, 3, 4, 5",
			args: args{
				input: []int{1, 2, 3, 4, 5},
			},
			want: 180,
		},
		{
			name: "100, 100, 2",
			args: args{
				input: []int{100, 100, 2},
			},
			want: 20000,
		},
		{
			name: "1, 1, 1",
			args: args{
				input: []int{1, 1, 1},
			},
			want: 3,
		},
		{
			name: "1, 10, 99",
			args: args{
				input: []int{1, 10, 99},
			},
			want: 11 * 99,
		},
		{
			name: "99, 10, 1",
			args: args{
				input: []int{1, 10, 99},
			},
			want: 11 * 99,
		},
		{
			name: "1, 99, 1",
			args: args{
				input: []int{1, 99, 1},
			},
			want: 101,
		},
		{
			name: "1, 2, 4, 8, 4, 2, 1",
			args: args{
				input: []int{1, 2, 4, 8, 4, 2, 1},
			},
			want: 12 * 12 * 8,
		},
		{
			name: "1000, 1000, 1, 1000",
			args: args{
				input: []int{1000, 1000, 1, 1000},
			},
			want: 1000 * 1001 * 1000,
		},
		{
			name: "1, 2, 3, 4, 5, 6, 7, 8, 9, 10",
			args: args{
				input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 3 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveMaxNumberWithSignsRecursive(tt.args.input); got != tt.want {
				t.Errorf("SolveMaxNumberWithSignsRecursive() = %v, want %v", got, tt.want)
			}
			if got := SolveMaxNumberWithSignsDP(tt.args.input); got != tt.want {
				t.Errorf("SolveMaxNumberWithSignsDP() = %v, want %v", got, tt.want)
			}

		})
	}
}
