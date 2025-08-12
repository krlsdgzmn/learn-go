package main

import (
	"fmt"
)

func main() {
	// if else
	age := 18
	if age >= 18 {
		fmt.Println("You're an adult.")
	} else if age >= 13 {
		fmt.Println("You're a teen.")
	} else {
		fmt.Println("You're a child.")
	}

	// switch
	day := 4
	switch day {
	case 1:
		fmt.Println("It's Monday.")
	case 2:
		fmt.Println("It's Tuesday.")
	default:
		fmt.Println("IDK")
	}

	// for loop
	for i := 0; i < 3; i++ {
		fmt.Printf("for %d\n", i)
	}

	// modern for loop with range
	for i := range [3]int{} {
		fmt.Printf("range %d\n", i)
	}

	// while loop but with for keyword
	counter := 0
	for counter < 3 {
		fmt.Printf("while %d\n", counter)
		counter++
	}

	// infinite loop
	// for {
	// 	// if condition is met use break
	// }

	// arrays
	arr := [3]int{1, 2, 3}
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("last value: %v\n", arr[len(arr)-1])

	// slices
	slcs := arr[1:]
	slcs = append(slcs, 5)
	fmt.Printf("slcs: %v\n", slcs)
	for i, v := range slcs {
		fmt.Printf("index: %d value: %d\n", i, v)
	}

	// map
	students := map[int]string{
		20214: "John Karlos",
		20215: "Alyhanna Faith",
	}
	for k, v := range students {
		fmt.Printf("key: %d value: %s\n", k, v)
	}

	// check if key exists
	_, exists := students[20215]
	if exists {
		fmt.Println("does exists")
	} else {
		fmt.Println("does not exists")
	}

	// delete
	delete(students, 20215)
	for k, v := range students {
		fmt.Printf("key: %d value: %s\n", k, v)
	}
}
