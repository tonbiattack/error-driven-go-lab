package errorlearning

import "sort"

func RankedScores(scores []int) []int {
	result := append([]int(nil), scores...)
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}
