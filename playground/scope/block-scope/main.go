package main

import "fmt"

func print() { // block scope starts here
	var hello = "Hello, world!"
	const ok = true
	// hello and ok are only accessible in this function
	fmt.Println(hello, ok)
} // block scope ends here

func main() {
	// hello and ok are not accessible here
	// ERROR: hello and ok are not defined
	// fmt.Println(hello, ok)

	// call print() to get access to hello and ok
	print()
}
