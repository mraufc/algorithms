package dynamicprogramming

import "testing"

func TestSolveFibonacciRecursive(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "tc1",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "tc2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "tc3",
			args: args{n: 9},
			want: 34,
		},
		{
			name: "tc4",
			args: args{n: 17},
			want: 1597,
		},
		{
			name: "tc5",
			args: args{n: 24},
			want: 46368,
		},
		// {
		// 	name: "tc6",
		// 	args: args{n: 79},
		// 	want: 23416728348467685,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveFibonacciRecursive(tt.args.n); got != tt.want {
				t.Errorf("SolveFibonacciRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveFibonacciDP(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "tc1",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "tc2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "tc3",
			args: args{n: 9},
			want: 34,
		},
		{
			name: "tc4",
			args: args{n: 17},
			want: 1597,
		},
		{
			name: "tc5",
			args: args{n: 24},
			want: 46368,
		},
		{
			name: "tc6",
			args: args{n: 80},
			want: 23416728348467685,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r1, r2 := SolveFibonacciDP(tt.args.n), SolveFibonacciDPSO(tt.args.n); r1 != tt.want || r2 != tt.want {
				t.Errorf("SolveFibonacciDP() = %v, SolveFibonacciDPSO() = %v want %v", r1, r2, tt.want)
			}
		})
	}
}
