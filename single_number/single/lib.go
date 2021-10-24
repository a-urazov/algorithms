package single

func Int(a []int) int {
	mask := 0
	for _, i := range a {
		mask ^= i
	}
	return mask
}
