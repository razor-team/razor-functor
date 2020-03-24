package functor

import (
	"razor-functor/src/dynamic"
	"testing"
)

func TestCall(t *testing.T) {
	result := dynamic.Call(func(i int) int {
		return i + 1
	}, 1, 2, 3, 4)
	lenResult, lenTest := len(result), 4
	if lenResult != lenTest {
		t.Errorf("count of test data is not equal to count of function reutrned result\n")
		t.Errorf("expected is %v, but actual is %v\n", lenResult, lenTest)
		return
	}
	c := 1
	for j, i := range result {
		c = c + 1
		v := i.Interface()
		if c != v {
			t.Errorf("value %d: expected is %v, but actual is %v\n", j, v, c)
		}
	}
}
