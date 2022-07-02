package main

// file scope
import "fmt"

// package-level declarations
const ok = true

func main() {
	var hello = "Hello, world!"
	// hello and ok are accessible in this function
	fmt.Println(hello, ok)
}
