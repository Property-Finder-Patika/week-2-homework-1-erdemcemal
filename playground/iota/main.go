package main

import "fmt"

// a, b, c, d are autoincrementing iota
const (
	a = iota
	b
	c
	d
)

// BitShiftingWithIota kb, mb, and gb bit shifting with  using iota
func BitShiftingWithIota() {
	const (
		_  = iota
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
	)
	fmt.Println(kb, mb, gb) // "1024 1048576 1073741824"
}

// BitShiftingWithoutUsingIota shift bits without using iota
func BitShiftingWithoutUsingIota() {
	const (
		kb = 1024
		mb = 1024 * kb
		gb = 1024 * mb
	)
	fmt.Println(kb, mb, gb) // "1024 1048576 1073741824"
}

func main() {
	// simple autoincrementing iota
	fmt.Println(a, b, c, d) // "0 1 2 3"

	// bit shifting with iota
	BitShiftingWithIota() // "1024 1048576 1073741824"

	// bit shifting without using iota
	BitShiftingWithoutUsingIota() // "1024 1048576 1073741824"
}
