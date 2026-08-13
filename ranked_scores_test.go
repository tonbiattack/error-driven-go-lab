package errorlearning

import (
	"reflect"
	"testing"
)

func TestG001RankedScoresDoesNotMutateInput(t *testing.T) {
	source := []int{30, 10, 20}

	got := RankedScores(source)

	if !reflect.DeepEqual(got, []int{30, 20, 10}) {
		t.Fatalf("RankedScores() = %v, want [30 20 10]", got)
	}
	if !reflect.DeepEqual(source, []int{30, 10, 20}) {
		t.Fatalf("source = %v, want [30 10 20]", source)
	}
}
