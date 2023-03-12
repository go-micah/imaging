package utils

import "testing"

func TestAbs(t *testing.T) {
	want := 100
	got_pos := Abs(100)
	got_neg := Abs(-100)

	if (want != got_pos) || (want != got_neg) {
		t.Errorf("expected %q, got %q and %q", want, got_pos, got_neg)
	}
}
