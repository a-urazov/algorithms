package array

import "testing"

func TestMaxSubarraySum(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"max sub array sum",
			args{[]int{1, 2, 3, 1, -1, -10, 9, 5}},
			14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubarraySum(tt.args.a); got != tt.want {
				t.Errorf("MaxSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
