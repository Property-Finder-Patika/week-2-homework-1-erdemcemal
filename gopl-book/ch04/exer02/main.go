package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var modeFlag = flag.String("mode", "sha256", "hash: sha256, sha384, or sha512")

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Printf("Error reading from stdin: %v\n", err)
		os.Exit(1)
	}
	printHash(bytes, *modeFlag)
}

// Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.
func printHash(bs []byte, mode string) {
	switch mode {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256(bs))
		break
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384(bs))
		break
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512(bs))
		break
	default:
		fmt.Printf("Invalid mode: %s\n", mode)
		os.Exit(1)
	}
}
