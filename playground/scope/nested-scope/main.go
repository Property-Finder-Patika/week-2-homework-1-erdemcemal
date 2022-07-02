package main

import "fmt"

var v = "outer"

// this example shows how to use the package-level variable v and how to
// use the scope of a function.
func main() {
	fmt.Println(v) // outer
	{
		v := "inner"
		fmt.Println(v) // inner
		{
			v := "innermost"
			fmt.Println(v) // innermost
		}
	}
	{
		fmt.Println(v) // outer
	}
	fmt.Println(v) // outer
}
