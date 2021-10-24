package array

func MaxSubarraySum(a []int) int {
	current := a[0]
	maximum := a[0]
	for i := 1; i < len(a); i++ {
		n := a[i]
		current = maxBetweenInts(current + n, n)
		maximum = maxBetweenInts(current, maximum)
	}
	return maximum
}

func maxBetweenInts(i, n int) int {
	if i > n {
		return i
	}
	return n
}