package disappeared

import (
	"reflect"
	"testing"
)

func TestInts(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"With disappeared elements [5, 6]",
			args{[]int{1, 2, 3, 4, 2, 3, 7, 8, 9}},
			[]int{5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ints(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ints() = %v, want %v", got, tt.want)
			}
		})
	}
}
