package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "ch04/exer08: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsDigit(r) {
			counts["Digit"]++
		}
		if unicode.IsLetter(r) {
			counts["Letter"]++
		}
		if unicode.IsLower(r) {
			counts["Lower"]++
		}
		if unicode.IsUpper(r) {
			counts["Upper"]++
		}
		if unicode.IsSpace(r) {
			counts["Space"]++
		}
		if unicode.IsSymbol(r) {
			counts["Symbol"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for r, n := range counts {
		fmt.Printf("%s\t%d\n", r, n)
	}
	fmt.Printf("Invalid: %d\n", invalid)
}
