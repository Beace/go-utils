package hello

import "testing"

func TestSum(t *testing.T) {
	want := 3

	if got := Sum(1, 2); got != want {
		t.Errorf("Sum() = %q, want %q", got, want)
	}
}
