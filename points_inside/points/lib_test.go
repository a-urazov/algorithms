package points

import "testing"

func TestInside(t *testing.T) {
	type args struct {
		t     [][]int
		q     []int
		depth int
		i     int
		j     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test is area has points [i..j] inside them",
			args{
				[][]int{
					{1, 2, 3},
					{1, 2, 3},
					{1, 2, 3},
				},
				[]int{1, 2, 3},
				2,
				0,
				1,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInside(tt.args.t, tt.args.q, tt.args.depth, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Inside() = %v, want %v", got, tt.want)
			}
		})
	}
}
