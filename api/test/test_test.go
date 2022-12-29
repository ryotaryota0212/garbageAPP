package test

import "testing"

func TestSum(t *testing.T) {
	s := Sum([]int{2, 3, 4, 5})
	if s != 14 {
		t.Error("error")
	}
}
