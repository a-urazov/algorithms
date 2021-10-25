package linkedlist

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func HasCycle(list *LinkedList) bool {
	if list == nil {
		return false
	}
	fast := list
	slow := fast
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
