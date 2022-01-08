package main

import (
	"math"
	"testing"
)

func isCloseEnough(a, b float64) bool {
	const acc = 0.001
	return (math.Abs(a-b) < acc)
}

func TestCiecl(t *testing.T) {
	s := 56.00

	expect := 8.445

	val := Circle(s)

	if !isCloseEnough(val, expect) {
		t.Errorf("expected %.3f got %.3f", expect, val)
	}
}
