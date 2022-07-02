package main

import "testing"

func TestDifferentBits(t *testing.T) {
	var tests = []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"a", "a", 0},
		{"a", "b", 126},
	}
	for _, tt := range tests {
		if got := sha256PopCount(tt.a, tt.b); got != tt.want {
			t.Errorf("sha256PopCount(%q, %q) = %v", tt.a, tt.b, got)
		}
	}

}
