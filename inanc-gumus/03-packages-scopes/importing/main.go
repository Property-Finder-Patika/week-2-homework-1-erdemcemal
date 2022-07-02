package main

import "fmt"
import f_second "fmt"
import fm_third "fmt"

// ---------------------------------------------------------
// EXERCISE: Rename imports
//
//  1- Import fmt package three times with different names
//
//  2- Print a few messages using those imports
//
// EXPECTED OUTPUT
//  hello
//  hey
//  hi
// ---------------------------------------------------------

// ?
// ?
// ?

func main() {
	fmt.Println("hello")
	f_second.Println("hey")
	fm_third.Println("hi")
}
