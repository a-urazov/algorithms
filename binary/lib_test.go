package binary

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		a []int
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test binary search",
			args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
