package utils

import "testing"

func TestAbs(t *testing.T) {
	want := 100
	got := Abs(-100)

	if want != got {
		t.Errorf("expected %q, got %q", want, got)
	}
}
