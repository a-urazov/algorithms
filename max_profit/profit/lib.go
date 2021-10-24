package profit

func Int(a []int) int {
	profit := 0
	minimal := a[0]
	for i := 1; i < len(a); i++ {
		price := a[i]
		profit = maxBetweenInts(profit, price - minimal)
		minimal = minBetweenInts(minimal, price)
	}
	return profit
}

func minBetweenInts(minimal, price int) int {
	if minimal < price {
		return minimal
	}
	return price
}

func maxBetweenInts(profit, i int) int {
	if profit > i {
		return profit
	}
	return i
}
