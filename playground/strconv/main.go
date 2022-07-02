package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ConvertStrToInt("123"))
	fmt.Println(ConvertIntToStr(123))
}

// ConvertStrToInt a string to an integer
func ConvertStrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// ConvertIntToStr an integer to a string
func ConvertIntToStr(i int) string {
	str := strconv.Itoa(i)
	return str
}
