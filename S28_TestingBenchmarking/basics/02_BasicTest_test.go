package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{20, 22}, 42},
		test{[]int{3, 22}, 25},
		test{[]int{4, 38}, 42},
		test{[]int{10, 12}, 22},
	}

	for _, v := range tests {
		if mySum(v.data...) != v.answer {
			t.Error("Expected", v.answer, "Got", mySum(v.data...))
		}
	}
}
