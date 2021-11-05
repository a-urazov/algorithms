package points

func IsInside(area [][]int, location []int, depth, from, to int) int {
	if from == to {
		return 0
	}
	if to-from == 1 {
		return is_inside(area[from], location)
	}
	k := 1 - depth&1
	m := from + (to-from)/2
	diff := area[m][k] - location[k]
	if diff < location[2] {
		return IsInside(area, location, depth+1, from, m)
	}
	if diff > -location[2] {
		return IsInside(area, location, depth+1, m+1, to)
	}

	return IsInside(area, location, depth+1, from, m) + is_inside(area[m], location) + IsInside(area, location, depth+1, m+1, to)
}

func is_inside(t, q []int) int {
	for _, vt := range q {
		for _, vq := range t {
			if vt == vq {
				return 1
			}
		}
	}

	return 0
}
