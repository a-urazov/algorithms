package disappeared

func Ints(a []int) []int {
	i := 0
	for i < len(a) {
		p := a[i] - 1 // correct position
		if a[i] != a[p] {
			a[i], a[p] = a[p], a[i]
		} else {
			i += 1
		}
	}
	missing := make([]int, 0)
	for i := 0; i < len(a); i++ {
		if a[i] != i+1 {
			missing = append(missing, i+1)
		}
	}
	return missing
}
