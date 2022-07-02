package main

import "fmt"

var apples int32 = 1
var oranges int16 = 2

func main() {
	fmt.Println(Compote())

	// floating point conversion change the value or lose precision
	f := 3.141 // a float64
	i := int(f)
	fmt.Println(f, i) // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"

	// type conversion is not always safe
	var myInt int = 65
	var myUint uint = uint(myInt)
	var myFloat float64 = float64(myInt)
	fmt.Println(myInt, myUint, myFloat) // "65 65 65.0"
}

func Compote() int {
	// var x int = apples + oranges // compile error: Invalid operation: apples + oranges (mismatched types int32 and int16)

	// explicit type conversion
	var x int = int(apples) + int(oranges)
	return x
}
