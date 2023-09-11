package main

import "testing"

func TestSum(t *testing.T){
	result := sum(2,5)

	if result != 7 {
		t.Error("The result must be 7")
	}
}