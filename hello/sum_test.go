package hello

import "testing"

func TestSum(t *testing.T) {
	want := 3

	if got := Sum(1, 2); got != want {
		t.Errorf("Sum() = %q, want %q", got, want)
	}

	/*if got := Sum(2,1); got != 1 {
		t.Errorf("Sum() = %q, want %q", got, want)
	}*/
}


func TestPlus(t *testing.T) {
	want := 3

	if got := Minus(4, 1); got != want {
		t.Errorf("Sum() = %q, want %q", got, want)
	}
}