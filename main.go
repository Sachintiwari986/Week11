package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Addition:", Add(2, 3))
	fmt.Println("Subtration:", Subtract(2, 3))
	fmt.Println("Multiply:", Multiply(2, 3))

}
