package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "John"}
	// person := struct{Name string}{Name: "John"}

	fmt.Printf("person = %s\n", person.Name)

	// pointers
	fmt.Println("\nWithout Pointer")
	fmt.Printf("before: %s\n", person.Name)
	changeName(person, "Karlos")
	fmt.Printf("after: %s\n", person.Name)

	fmt.Println("\nWith Pointer")
	fmt.Printf("before: %s\n", person.Name)
	changeNameByPointer(&person, "Karlos")
	fmt.Printf("after: %s\n", person.Name)
}

func changeName(p Person, name string) {
	p.Name = name
	fmt.Printf("reassigning the value inside this function scope: %s\n", p.Name)
}

func changeNameByPointer(p *Person, name string) {
	p.Name = name
	fmt.Printf("reassigning the value inside this function scope: %s\n", p.Name)
}
