package main_test

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T)  {
	fmt.Println("test")
	test := 20
	if test != 20 {
		t.Errorf("not equal 20 %v", test)
	}
}