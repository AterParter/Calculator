package main

import "fmt"

func main() {
	var a int
	var b string
	var c int

	fmt.Println("Enter your First Digit : ")
	fmt.Scan(&a)
	fmt.Println("Enter your Opperator (plus, minus, mul, div) : ")
	fmt.Scan(&b)
	fmt.Println("Enter your Second Digit : ")
	fmt.Scan(&c)

	if b == "plus" {
		fmt.Println("-------------------")
		fmt.Println(a + c)
	}

	if b == "minus" {
		fmt.Println("-------------------")
		fmt.Println(a - c)
	}

	if b == "mul" {
		fmt.Println("-------------------")
		fmt.Println(a * c)
	}

	if b == "div" {
		fmt.Println("-------------------")
		fmt.Println(a / c)
	}

}
