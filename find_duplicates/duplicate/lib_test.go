package duplicate

import "testing"

func TestInts(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test array without duplicates",
			args{[]int{1, 2, 3}},
			false,
		},
		{
			"Test array with duplicates",
			args{[]int{1, 1, 2, 3}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ints(tt.args.a); got != tt.want {
				t.Errorf("Ints() = %v, want %v", got, tt.want)
			}
		})
	}
}
