package functor

import (
	"razor-functor/src/stl"
	"testing"
)

func TestFlatten(t *testing.T) {
	r := flatten(stl.List(1, 2, stl.List(1, 2, 3)))
	n := len(r)
	if n != 5 {
		t.Errorf("expected 5, but acutal is %d", n)
	}
}
