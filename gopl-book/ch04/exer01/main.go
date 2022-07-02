package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
func main() {
	fmt.Println(sha256PopCount("a", "b")) // 1
	fmt.Println(sha256PopCount("a", "a")) // 0
}

// Exercise 4.1:Write a function that counts the number of bits that are different in two SHA256 hashes
func sha256PopCount(a, b string) int {
	digestA := sha256.Sum256([]byte(a))
	digestB := sha256.Sum256([]byte(b))
	return popCount(digestA, digestB)
}

// popCount returns the number of bits set in x.
func popCount(a, b [32]byte) int {
	pop := 0
	for i := range a {
		pop += int(pc[a[i]^b[i]])
	}
	return pop
}
