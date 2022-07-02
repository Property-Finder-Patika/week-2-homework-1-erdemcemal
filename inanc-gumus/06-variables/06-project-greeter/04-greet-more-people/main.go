package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("There are not enough people.")
		os.Exit(1)
	}
	fmt.Printf("There are %d people!\n", len(args))
	for _, arg := range args {
		fmt.Printf("Hello great %s !\n", arg)
	}
}
