package main

import (
	"fmt"
)

func main() {
	// variables
	var name string = "JK"
	fmt.Printf("My name is %s. ", name)

	age := 23
	fmt.Printf("I am %d years old.\n", age)

	var street, city = "Rose Street", "Caloocan City"
	fmt.Printf("I lived in %s, %s\n", street, city)

	var isEmployed bool = true
	fmt.Printf("Is employed: %t\n", isEmployed)

	var (
		position string = "Software Engineer"
		salary   int    = 40000
	)
	fmt.Printf("I am a %s and my salary is %d\n", position, salary)

	// zero values
	var defInt int
	var defFloat float64
	var defString string
	var defBool bool

	fmt.Printf("int: %d float: %f string: %s bool: %t\n", defInt, defFloat, defString, defBool)

	// constants
	const typedPi float64 = 3.14
	const untypedPi = 3.14

	const (
		monday        = 1
		tuesday       = 2
		wednesday int = 3
	)

	const (
		jan = iota + 1 // 1
		feb            // 2
		mar            // 3
	)
	fmt.Printf("jan: %d feb: %d mar: %d\n", jan, feb, mar)

	// private var
	// camelCase := 1
	// public var
	// PascalCase := 1

	result := add(1, 3)
	fmt.Printf("result: %d\n", result)

	sum, prod := calcSumProd(1, 3)
	fmt.Printf("sum: %d prod: %d\n", sum, prod)
}

func add(a int, b int) int {
	return a + b
}

func calcSumProd(a, b int) (int, int) {
	return a + b, a * b
}
