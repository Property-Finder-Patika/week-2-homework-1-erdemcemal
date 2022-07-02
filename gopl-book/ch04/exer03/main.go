package main

import "fmt"

const SIZE = 6

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	reverseArr(&a)
	fmt.Println(a) // [6,5,4,3,2,1]
}

// Exercise 4.3:  Rewrite reverse to use an array pointer instead of a slice.
func reverseArr(a *[SIZE]int) {
	for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
