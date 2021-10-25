package linkedlist

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func HasCycle(list *LinkedList) bool {
	if list == nil {
		return false
	}
	var fast, slow *LinkedList = list, list
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func FindMiddle(list *LinkedList) *LinkedList {
	var fast, slow *LinkedList = list, list
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func Reverse(list *LinkedList) *LinkedList {
	var prev *LinkedList
	current := list
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func IsPolindrome(list *LinkedList) bool {
	if list == nil {
		return false
	}
	middle := FindMiddle(list)
	reversed := Reverse(middle)
	for reversed != nil {
		if reversed.Value != list.Value {
			return false
		}
		reversed = reversed.Next
		list = list.Next
	}
	return true
}
