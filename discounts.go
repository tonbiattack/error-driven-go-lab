package errorlearning

import "fmt"

func DiscountFor(code string, discounts map[string]int) (int, error) {
	discount, ok := discounts[code]
	if !ok {
		return 0, fmt.Errorf("unknown discount code: %s", code)
	}
	return discount, nil
}
