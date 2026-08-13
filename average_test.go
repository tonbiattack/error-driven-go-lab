package errorlearning

import "testing"

func TestG003AverageKeepsFraction(t *testing.T) {
	if got := Average(5, 2); got != 2.5 {
		t.Fatalf("Average(5, 2) = %v, want 2.5", got)
	}
}

func TestG003AverageReturnsWholeNumberAsFloat(t *testing.T) {
	if got := Average(6, 2); got != 3.0 {
		t.Fatalf("Average(6, 2) = %v, want 3", got)
	}
}
