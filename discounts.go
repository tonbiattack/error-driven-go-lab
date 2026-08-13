package errorlearning

func DiscountFor(code string, discounts map[string]int) (int, error) {
	return discounts[code], nil
}
