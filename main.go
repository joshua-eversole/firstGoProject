package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("Test #2. please commit")
	var numerator int = 11
	var denominator int = 2
	var result, remainder int = intDivision(numerator, denominator)
	fmt.Println("The result of dividing", numerator, "by", denominator, "is", result, "with a remainder of", remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}
