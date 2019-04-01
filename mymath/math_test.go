package mymath

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2, 3) != 6 {
		t.Error("Add(1,2,3) not equal to 6.")
	}
}
