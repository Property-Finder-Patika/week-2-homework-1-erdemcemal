package main

import (
	"reflect"
	"testing"
)

func TestReverseArrayUsingPointer(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, test := range tests {
		reverse(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("reverseArrayUsingPointer(%v) = %v, want %v", test.input, test.input, test.want)
		}
	}
}
