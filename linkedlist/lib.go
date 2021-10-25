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
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}