package profit

import "testing"

func TestInt(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Has profit 8",
			args{[]int{7, 1, 5, 3, 6, 9}},
			8,
		},
		{
			"Has profit 6",
			args{[]int{7, 1, 5, 3, 6, 4}},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.a); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
