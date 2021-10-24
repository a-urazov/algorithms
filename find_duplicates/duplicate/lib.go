package duplicate

// Ints
func Ints(a []int) bool {
	m := make(map[int]bool)
	for i := range a {
		if _, exists := m[a[i]]; exists {
			return true
		}
		m[a[i]] = true
	}
	return false
}
