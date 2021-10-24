package bits

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"bit counts for 2 is [0, 1, 1]",
			args{2},
			[]int{0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
