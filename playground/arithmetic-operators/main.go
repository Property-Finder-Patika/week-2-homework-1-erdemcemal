package main

import (
	"fmt"
	"math"
)

func main() {

	// numeric type operations
	numericTypeOperations()

	// boolean type operations
	booleanTypeOperations()

	// complex type operations
	complexTypeOperations()

	// string type operations
	stringTypeOperations()
}

// simple numeric type operations
func numericTypeOperations() {
	var a, b = 4, 5
	var res1 = (a + b) * (a + b) / 2 // Arithmetic operations

	a++ // Increment a by 1

	b += 10 // Increment b by 10

	var res2 = a ^ b // Bitwise XOR

	var r = 3.5
	var res3 = math.Pi * r * r // Operations on floating-point type

	fmt.Printf("res1 : %v, res2 : %v, res3 : %v\n", res1, res2, res3) // "res1 : 100, res2 : 15, res3 : 22.5"
}

// simple complex type operations
func complexTypeOperations() {
	var a = 3 + 5i
	var b = 2 + 4i

	var res1 = a + b
	var res2 = a - b
	var res3 = a * b
	var res4 = a / b

	fmt.Println(res1, res2, res3, res4)
}

// simple boolean type operations
func booleanTypeOperations() {
	var a, b = true, false
	var res1 = a && b             // Logical AND
	var res2 = a || b             // Logical OR
	var res3 = !a                 // Logical NOT
	fmt.Println(res1, res2, res3) // "false true false"
}

// simple string type operations
func stringTypeOperations() {
	var a = "Hello"
	var b = "World"
	var res1 = a + " " + b  // Concatenation
	var res2 = a + b        // Concatenation
	fmt.Println(res1, res2) // "Hello World" "HelloWorld"
}
