package missing

import "errors"

func Int(a []int) (int, error) {
	n := len(a)
	result := n * (n + 1)/ 2 - sumInt(a)
	if result < a[0] || result > a[n-1] {
		return 0, errors.New("missing number not found")
	}
	return result, nil
}

func sumInt(a []int) int {
	n := 0
	for i := 0; i < len(a); i++ {
		n += a[i]
	}
	return n 
}
