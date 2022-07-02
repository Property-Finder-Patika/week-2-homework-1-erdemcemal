package main

func main() {
	name := "World"
	// main and print files belong to the same package
	// so they can access each other's variables
	// calling greetings() from print.go
	greetings(name)
}

// following function will throw an error
/*func greetings(name string) {
	fmt.Println("Hello, ", name)
}*/

// *******EXPLANATION********
// ERROR: 'greetings' redeclared in this package
// "this-block" means = "main package"
