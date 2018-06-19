package dynamicprogramming

import "testing"

func TestRodCutting(t *testing.T) {
	type args struct {
		prices []int
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				[]int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30},
				4,
			},
			want: 10,
		},
		{
			name: "tc2",
			args: args{
				[]int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30},
				6,
			},
			want: 17,
		},
		{
			name: "tc3",
			args: args{
				[]int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30},
				10,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveRodCuttingRecursiveTopDown(tt.args.prices, tt.args.length); got != tt.want {
				t.Errorf("SolveRodCuttingRecursiveTopDown() = %v, want %v", got, tt.want)
			}
			if got := SolveRodCuttingMemoized(tt.args.prices, tt.args.length); got != tt.want {
				t.Errorf("SolveRodCuttingMemoized() = %v, want %v", got, tt.want)
			}
			if got := SolveRodCuttingTabulated(tt.args.prices, tt.args.length); got != tt.want {
				t.Errorf("SolveRodCuttingTabulated() = %v, want %v", got, tt.want)
			}

		})
	}
}
