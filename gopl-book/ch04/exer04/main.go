package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	rotate(s, 3)
	fmt.Println(s) // [4, 5, 6, 1, 2, 3]
}

// Exercise 4.4: Write a version of rotate that operates in a single pass
func rotate(s []int, n int) {
	num := n % len(s)
	double := append(s, s[:num]...)
	copy(s, double[num:num+len(s)])
}
