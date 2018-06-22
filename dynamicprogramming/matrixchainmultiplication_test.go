package dynamicprogramming

import "testing"

func TestSolveMatrixMultiplicationRecursive(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{input: []int{40, 20, 30}},
			want: 24000,
		},
		{
			name: "tc2",
			args: args{input: []int{40, 20, 30, 10, 30}},
			want: 26000,
		},
		{
			name: "tc3",
			args: args{input: []int{10, 20, 30, 40, 30}},
			want: 30000,
		},
		{
			name: "tc4",
			args: args{input: []int{10, 20, 30}},
			want: 6000,
		},
		{
			name: "tc5",
			args: args{input: []int{30, 35, 15, 5, 10, 20, 25}},
			want: 15125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveMatrixMultiplicationRecursive(tt.args.input); got != tt.want {
				t.Errorf("SolveMatrixMultiplicationRecursive() = %v, want %v", got, tt.want)
			}
			if got := SolveMatrixMultiplicationDP(tt.args.input); got != tt.want {
				t.Errorf("SolveMatrixMultiplicationDP() = %v, want %v", got, tt.want)
			}

		})
	}
}
