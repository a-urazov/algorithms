package linkedlist

import (
	"reflect"
	"testing"
)

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

func TestFindMiddle(t *testing.T) {
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
	list5.Next = nil
	tests := []struct {
		name string
		args args
		want *LinkedList
	}{
		{
			"LinkedList middle",
			args{list1},
			list3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMiddle(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
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
	list5.Next = nil
	tests := []struct {
		name string
		args args
		want *LinkedList
	}{
		{
			"LinkedList reversed",
			args{list1},
			list5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPolindrome(t *testing.T) {
	type args struct {
		list *LinkedList
	}
	list1 := &LinkedList{1, nil}
	list2 := &LinkedList{2, nil}
	list3 := &LinkedList{2, nil}
	list4 := &LinkedList{1, nil}
	list1.Next = list2
	list2.Next = list3
	list3.Next = list4
	list4.Next = nil
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"LinkedList is polindrome",
			args{list1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPolindrome(tt.args.list); got != tt.want {
				t.Errorf("IsPolindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
