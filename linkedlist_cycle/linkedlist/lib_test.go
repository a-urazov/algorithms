package linkedlist

import "testing"

func TestHasCycle(t *testing.T) {
	type args struct {
		list *LinkedList
	}
	
	list1 := &LinkedList{1, nil}
	list2 := &LinkedList{2, nil}
	list3 := &LinkedList{3, nil}
	list4 := &LinkedList{4, nil}
	list5 := &LinkedList{5, nil}
	list1.Next = list2
	list2.Next = list3
	list3.Next = list4
	list4.Next = list5
	list5.Next = list1

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"LinkedList has cycles",
			args{list1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle(tt.args.list); got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
