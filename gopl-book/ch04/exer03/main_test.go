package main

import (
	"reflect"
	"testing"
)

func TestReverseArrayUsingPointer(t *testing.T) {
	var tests = []struct {
		input [SIZE]int
		want  [SIZE]int
	}{
		{[SIZE]int{1, 2, 3, 4, 5, 6}, [SIZE]int{6, 5, 4, 3, 2, 1}},
		{[SIZE]int{10, 20, 30, 40, 50, 60}, [SIZE]int{60, 50, 40, 30, 20, 10}},
	}
	for _, test := range tests {
		input := test.input
		want := test.want
		reverseArr(&input)
		if !reflect.DeepEqual(input, want) {
			t.Errorf("reverseArrayUsingPointer(%v) = %v, want %v", input, input, want)
		}
	}
}
