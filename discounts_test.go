package errorlearning

import "testing"

func TestG002DiscountForRejectsUnknownCode(t *testing.T) {
	_, err := DiscountFor("UNKNOWN", map[string]int{"WELCOME": 10})

	if err == nil {
		t.Fatal("DiscountFor() error = nil, want an error for an unknown code")
	}
}

func TestG002DiscountForReturnsRegisteredValue(t *testing.T) {
	got, err := DiscountFor("WELCOME", map[string]int{"WELCOME": 10})

	if err != nil {
		t.Fatalf("DiscountFor() error = %v, want nil", err)
	}
	if got != 10 {
		t.Fatalf("DiscountFor() = %d, want 10", got)
	}
}
