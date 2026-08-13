package errorlearning

import "sort"

func RankedScores(scores []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	return scores
}
