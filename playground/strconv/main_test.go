package main

import "testing"

// test the ConvertStrToInt function
func TestConvertStrToInt(t *testing.T) {
	type test struct {
		input string
		want  int
	}
	tests := []test{
		{"123", 123},
		{"-123", -123},
		{"0", 0},
		{"123a", 0},
	}
	for _, v := range tests {
		if got := ConvertStrToInt(v.input); got != v.want {
			t.Errorf("ConvertStrToInt(%q) = %v, want %v", v.input, got, v.want)
		}
	}
}

// test the TestConvertIntToStr function
func TestConvertIntToStr(t *testing.T) {
	type test struct {
		input int
		want  string
	}
	tests := []test{
		{123, "123"},
		{-123, "-123"},
		{0, "0"},
	}
	for _, v := range tests {
		if got := ConvertIntToStr(v.input); got != v.want {
			t.Errorf("ConvertIntToStr(%q) = %v, want %v", v.input, got, v.want)
		}
	}
}
