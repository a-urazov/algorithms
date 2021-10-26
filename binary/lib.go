package binary

func Search(a []int, v int) int {
	n := len(a)
	low, high := 0, n
	for low <= high {
		middle := low + (high-low)/2
		if a[middle] == v {
			return middle
		}
		if a[middle] > v {
			high = middle - 1
			continue
		}
		low = middle + 1
	}
	return -1
}
