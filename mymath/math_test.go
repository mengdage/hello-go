package mymath_test

import (
	"testing"

	"github.com/mengdage/hello-go/mymath"
)

func TestAdd(t *testing.T) {
	if mymath.Add(1, 2, 3) != 6 {
		t.Error("Add(1,2,3) not equal to 6.")
	}
}
